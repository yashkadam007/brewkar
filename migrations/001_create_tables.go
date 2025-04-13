package main

import (
	"fmt"
	"log"

	"github.com/yashkadam007/brewkar/internal/config"
	"github.com/yashkadam007/brewkar/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to the database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User,
		cfg.Database.Password, cfg.Database.DBName, cfg.Database.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate the schema
	fmt.Println("Creating database tables...")

	err = db.AutoMigrate(
		&domain.User{},
		// Add other models here as needed
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	fmt.Println("Migration completed successfully")
}
