package middlewares

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Validator is a global validator instance
var validate = validator.New()

// ValidatePhone custom validation for phone numbers
func ValidatePhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	// Simple regex for phone numbers (e.g., "09338668019")
	// You can adjust this regex to fit your phone format
	phoneRegex := `^(?:0|98|\+98|\+980|0098|098|00980)?(9\d{9})`
	matched, _ := regexp.MatchString(phoneRegex, phone)
	return matched
}

// ValidateDTO middleware function to validate the body of the incoming request
func ValidateDTO(dto interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the body into the given DTO struct
		if err := c.BodyParser(dto); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Failed to parse the request body",
			})
		}

		// Register the custom phone validation function
		validate.RegisterValidation("phone", ValidatePhone)

		// Validate the parsed DTO struct
		if err := validate.Struct(dto); err != nil {
			// Handle validation errors
			var validationErrors []string
			for _, err := range err.(validator.ValidationErrors) {
				validationErrors = append(validationErrors, err.Error())
			}

			log.Println("Validation errors:", validationErrors)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": validationErrors,
			})
		}

		// Continue with the request if validation is successful

		c.Locals("validatedDTO", dto)

		return c.Next()
	}
}
