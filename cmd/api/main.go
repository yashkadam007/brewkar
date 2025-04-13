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
	"github.com/yashkadam007/brewkar/internal/service"
	"github.com/yashkadam007/brewkar/pkg/logger"

	"github.com/gin-gonic/gin"
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

	// Set up Gin router
	router := gin.Default()

	// Register ping endpoint for health checks
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "pong",
		})
	})

	// Register routes
	v1 := router.Group("/v1")
	{
		// Auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
			auth.POST("/refresh", authController.Refresh)
		}

		// Protected routes
		// api := v1.Group("")
		// TODO: Add auth middleware

		// User routes
		// users := api.Group("/users")
		// {
		// 	// users.GET("/me", userController.GetProfile)
		// 	// users.PUT("/me", userController.UpdateProfile)
		// }

		// 	// Bean routes
		// 	beans := api.Group("/beans")
		// 	{
		// 		beans.GET("", beanController.GetAll)
		// 		beans.POST("", beanController.Create)
		// 		beans.GET("/:id", beanController.GetByID)
		// 		beans.PUT("/:id", beanController.Update)
		// 		beans.DELETE("/:id", beanController.Delete)
		// 		beans.POST("/:id/images", beanController.UploadImages)
		// 	}

		// 	// Recipe routes
		// 	recipes := api.Group("/recipes")
		// 	{
		// 		recipes.GET("", recipeController.GetAll)
		// 		recipes.POST("", recipeController.Create)
		// 		recipes.GET("/:id", recipeController.GetByID)
		// 		recipes.PUT("/:id", recipeController.Update)
		// 		recipes.DELETE("/:id", recipeController.Delete)
		// 		recipes.POST("/:id/images", recipeController.UploadImages)
		// 		recipes.GET("/public", recipeController.GetPublic)
		// 	}

		// 	// Brew log routes
		// 	brewLogs := api.Group("/brew-logs")
		// 	{
		// 		brewLogs.GET("", brewLogController.GetAll)
		// 		brewLogs.POST("", brewLogController.Create)
		// 		brewLogs.GET("/:id", brewLogController.GetByID)
		// 		brewLogs.PUT("/:id", brewLogController.Update)
		// 		brewLogs.DELETE("/:id", brewLogController.Delete)
		// 		brewLogs.POST("/:id/images", brewLogController.UploadImages)
		// 	}

		// 	// TODO: Add other routes (analytics, social, etc.)
	}

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
