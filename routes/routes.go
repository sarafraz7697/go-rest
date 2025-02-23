package routes

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes initializes all the API routes for the application.
func SetupRoutes(app *fiber.App) {
	// Create a group for API routes
	api := app.Group("/api")
	// Create a version 1 group under the API routes
	v1 := api.Group("/v1")

	// Setup authentication routes
	SetupAuthRoute(app, v1)

	// Setup product-related routes
	SetupProductRoutes(app, v1)
}
