package database

import (
	"fmt"
	"log"
	"rest/config"
	models "rest/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect initializes the database connection
func Connect() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.AppConfig.DBHost,
		config.AppConfig.DBUser,
		config.AppConfig.DBPass,
		config.AppConfig.DBName,
		config.AppConfig.DBPort,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Database connected successfully")

	// Automatically migrate all models (tables)
	DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Cart{}, &models.Order{})
}
