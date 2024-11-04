package routes

import (
	"rest/models"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
	app.Post("/products", func(c *fiber.Ctx) error {
		product := new(models.Product)
		if err := c.BodyParser(product); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}
		// product.BeforeSave() // Set timestamps

		// Here, you would typically save the product to the database

		return c.Status(fiber.StatusCreated).JSON(product)
	})
}
