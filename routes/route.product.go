package routes

import (
	"rest/controllers"
	. "rest/middlewares"
	"rest/repository"
	. "rest/routes/dtos/product"
	"rest/services"

	"github.com/gofiber/fiber/v2"
)

// SetupProductRoutes initializes product-related routes
func SetupProductRoutes(app *fiber.App, router fiber.Router) {
	api := router.Group("/product", AuthMiddleware())

	// Initialize core components (Repo, Service, Controller)
	productRepo := repository.NewProductRepository()
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	// Define routes
	api.Get("/", productController.GetAllProducts)    // Get all products
	api.Get("/:id", productController.GetProductByID) // Get product by ID

	api.Post("/", RoleMiddleware(ADMIN), ValidateDTO(&CreateProductDTO{}), productController.CreateProduct)   // Create product
	api.Put("/:id", RoleMiddleware(ADMIN), ValidateDTO(&UpdateProductDTO{}), productController.UpdateProduct) // Update product
	api.Delete("/:id", RoleMiddleware(ADMIN), productController.DeleteProduct)                                // Delete product
}
