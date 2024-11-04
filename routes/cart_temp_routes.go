package routes

import (
	"rest/models"

	"github.com/gofiber/fiber/v2"
)

func CartTempRoutes(app *fiber.App) {
	app.Post("/cart", func(c *fiber.Ctx) error {
		cartItem := new(models.CartTemp)
		if err := c.BodyParser(cartItem); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}
		// cartItem.BeforeSave() // Set timestamps

		// Here, you would typically save the cart item to the database

		return c.Status(fiber.StatusCreated).JSON(cartItem)
	})
}
