package middlewares

import (
	"rest/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware verifies JWT and extracts user role
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get token from Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
		}

		// Extract token part
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
		}

		token := tokenParts[1]

		// Validate and parse token
		claims, err := utils.ParseJWT(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		// Store user ID and role in Locals
		c.Locals("userId", claims["userId"])
		c.Locals("role", claims["role"])

		return c.Next()
	}
}
