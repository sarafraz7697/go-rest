// main.go
package main

import (
	"fmt"
	"os"
	"rest/routes"

	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Access the environment variables
	port := os.Getenv("PORT")
	// dbUser := os.Getenv("DB_USER")
	// dbPass := os.Getenv("DB_PASS")
	// secretKey := os.Getenv("SECRET_KEY")

	app := fiber.New()

	// Register routes
	routes.UserRoutes(app)
	routes.CategoryRoutes(app)
	routes.ProductRoutes(app)
	routes.CartTempRoutes(app)
	routes.OrdersRoutes(app)

	// Start server
	app.Listen(fmt.Sprintf(":%s", port))
}
