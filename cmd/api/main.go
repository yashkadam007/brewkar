package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yashkadam007/brewkar/internal/config"
	"github.com/yashkadam007/brewkar/internal/controller"
	"github.com/yashkadam007/brewkar/internal/repository"
	"github.com/yashkadam007/brewkar/internal/router"
	"github.com/yashkadam007/brewkar/internal/service"
	"github.com/yashkadam007/brewkar/pkg/logger"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	l := logger.NewLogger()
	defer l.Sync()

	// Connect to PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User,
		cfg.Database.Password, cfg.Database.DBName, cfg.Database.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		l.Fatal("Failed to connect to database", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		l.Fatal("Failed to get database connection", err)
	}
	defer sqlDB.Close()

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	ctx := context.Background()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		l.Warn("Failed to connect to Redis", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	// beanRepo := repository.NewBeanRepository(db)
	// recipeRepo := repository.NewRecipeRepository(db)
	// brewLogRepo := repository.NewBrewLogRepository(db)

	// Initialize services
	authService := service.NewAuthService(userRepo, cfg.JWT)
	// userService := service.NewUserService(userRepo)
	// beanService := service.NewBeanService(beanRepo)
	// recipeService := service.NewRecipeService(recipeRepo)
	// brewLogService := service.NewBrewLogService(brewLogRepo, beanRepo, recipeRepo)

	// Initialize controllers
	authController := controller.NewAuthController(authService)
	// userController := controller.NewUserController(userService)
	// beanController := controller.NewBeanController(beanService)
	// recipeController := controller.NewRecipeController(recipeService)
	// brewLogController := controller.NewBrewLogController(brewLogService)

	// Setup router with all routes
	router := router.SetupRouter(authController)

	// Start server
	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
	}

	// Graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Fatal("Failed to start server", err)
		}
	}()

	l.Info(fmt.Sprintf("Server started on port %s", cfg.Server.Port))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	l.Info("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		l.Fatal("Server forced to shutdown", err)
	}

	l.Info("Server exited")
}
