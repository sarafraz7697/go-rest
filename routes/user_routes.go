package routes

import (
	"rest/models"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Post("/users", func(c *fiber.Ctx) error {
		// dto := ValidateDto(c, CreateUserDto)
		user := new(models.User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}

		return c.Status(fiber.StatusCreated).JSON(user)
	})
}
