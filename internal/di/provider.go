package di

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/yashkadam007/brewkar/internal/config"
	"github.com/yashkadam007/brewkar/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ProvideConfig loads the application configuration
func ProvideConfig() (*config.Config, error) {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// ProvideLogger initializes the application logger
func ProvideLogger() *logger.Logger {
	return logger.NewLogger()
}

// ProvideDatabase initializes and returns the database connection
func ProvideDatabase(cfg *config.Config, l *logger.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User,
		cfg.Database.Password, cfg.Database.DBName, cfg.Database.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		l.Fatal("Failed to connect to database", err)
		return nil, err
	}

	return db, nil
}

// ProvideRedisClient initializes and returns the Redis client
func ProvideRedisClient(cfg *config.Config, l *logger.Logger) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	ctx := context.Background()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		l.Warn("Failed to connect to Redis", err)
		return nil, err
	}

	return rdb, nil
}
