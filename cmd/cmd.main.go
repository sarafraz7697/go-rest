package main

import (
	"fmt"
	"rest/config"
	"rest/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadConfig()

	// Initialize MongoDB
	config.ConnectDB()

	app := fiber.New()

	routes.SetupRoutes(app)

	// Start server
	port := config.GetEnvOrFatal("PORT")

	app.Listen(fmt.Sprintf(":%s", port))
}
