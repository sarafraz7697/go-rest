package routes

import (
	"rest/models"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(app *fiber.App) {
	app.Post("/categories", func(c *fiber.Ctx) error {
		category := new(models.Category)
		if err := c.BodyParser(category); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}
		// category.BeforeSave() // Set timestamps

		// Here, you would typically save the category to the database

		return c.Status(fiber.StatusCreated).JSON(category)
	})
}
