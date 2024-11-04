package routes

import (
	"rest/models"

	"github.com/gofiber/fiber/v2"
)

func OrdersRoutes(app *fiber.App) {
	app.Post("/orders", func(c *fiber.Ctx) error {
		order := new(models.Orders)
		if err := c.BodyParser(order); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}
		// order.BeforeSave() // Set timestamps

		// Here, you would typically save the order to the database

		return c.Status(fiber.StatusCreated).JSON(order)
	})
}
