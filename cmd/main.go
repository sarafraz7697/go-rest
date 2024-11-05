package main

import (
	"fmt"
	config "rest/config"
	"rest/internal/database"
	"rest/internal/handlers"
	"rest/internal/repository"
	"rest/internal/routes"

	services "rest/internal/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadConfig()
	database.Connect()

	app := fiber.New()

	// Initialize repositories and services
	userRepo := repository.NewUserRepository(database.DB)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)
	routes.SetupAuthRoutes(app, authHandler)

	// Start server
	port := config.GetEnvOrFatal("PORT")
	app.Listen(fmt.Sprintf(":%s", port))
}
