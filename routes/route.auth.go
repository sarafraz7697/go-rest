package routes

import (
	"rest/controllers"
	. "rest/middlewares"
	"rest/repository"
	"rest/services"

	. "rest/routes/dtos/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoute(app *fiber.App, router fiber.Router) {
	api := router.Group("/auth")
	{
		// Initiate repository, service, controller
		authRepo := repository.NewAuthRepository()
		authService := services.NewAuthService(authRepo)
		authController := controllers.NewAuthController(authService)

		// Define routes
		api.Post("/login", ValidateDTO(&LoginDTO{}), authController.Login)
		api.Post("/register", ValidateDTO(&RegisterDTO{}), authController.Register)
		api.Put("/profile", AuthMiddleware(), ValidateDTO(&UpdateProfileDTO{}), authController.UpdateProfile)
	}
}
