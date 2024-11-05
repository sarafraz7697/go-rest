package main

import (
	"fmt"
	"rest/config"
	"rest/internal/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize database connection
	database.Connect()

	app := fiber.New()

	// Start server
	port := config.AppConfig.Port
	app.Listen(fmt.Sprintf(":%s", port))
}
