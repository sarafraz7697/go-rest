package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Define roles as constants
const (
	ADMIN = "admin"
	USER  = "user"
)

// RoleMiddleware checks if the user has the required role to access a route
func RoleMiddleware(requiredRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the role from the request (from JWT, for example)
		userRole := c.Locals("role") // Assume we store the role in Locals after JWT validation

		// Ensure role exists
		if userRole == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized, no role found",
			})
		}

		// Convert interface{} to string
		roleStr := userRole.(string)

		// Check if the user has one of the required roles
		for _, role := range requiredRoles {
			if strings.EqualFold(roleStr, role) {
				return c.Next() // Proceed if role matches
			}
		}

		// If role doesn't match, return Forbidden
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden, insufficient permissions",
		})
	}
}
