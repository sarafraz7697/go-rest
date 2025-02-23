package utils

import "github.com/gofiber/fiber/v2"

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// SendResponse is a helper function to send a consistent JSON response.
func SendResponse(c *fiber.Ctx, statusCode int, data interface{}, message string) error {
	response := Response{
		Data:    data,
		Message: message,
	}

	return c.Status(statusCode).JSON(response)
}
