package routes

import (
	"rest/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App, authHandler *handlers.AuthHandler) {
	app.Post("/signin", authHandler.SignIn)
	// Add other auth-related routes like signup, etc.
}
