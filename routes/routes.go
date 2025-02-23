package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	SetupAuthRoute(app, v1)
	SetupProductRoutes(app, v1)

}
