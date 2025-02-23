package main

import (
	"fmt"
	"rest/config"
	"rest/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Connect to MongoDB
	config.ConnectDB()

	// Create a new Fiber application
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Start the server
	port := config.GetEnvOrFatal("PORT")
	app.Listen(fmt.Sprintf(":%s", port))
}
