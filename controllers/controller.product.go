package controllers

import (
	"context"
	"rest/models"
	dtos "rest/routes/dtos/product"
	"rest/services"
	"rest/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type ProductController struct {
	service *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{service: service}
}

// CreateProduct Handler
func (c *ProductController) CreateProduct(ctx *fiber.Ctx) error {
	// Get validated DTO from the Local storage
	dto, _ := ctx.Locals("validatedDTO").(*dtos.CreateProductDTO)

	// Convert DTO to Model
	product := &models.Product{}

	// Reflect Model from DTO
	utils.MapStruct(dto, product)

	newProduct, err := c.service.CreateProduct(context.Background(), product)
	if err != nil {
		return utils.SendResponse(ctx, fiber.StatusInternalServerError, nil, err.Error())
	}
	return ctx.Status(fiber.StatusCreated).JSON(newProduct)
}

// GetAllProducts Handler
func (c *ProductController) GetAllProducts(ctx *fiber.Ctx) error {
	products, err := c.service.GetAllProducts(context.Background())
	if err != nil {
		return utils.SendResponse(ctx, fiber.StatusInternalServerError, nil, err.Error())
	}
	return utils.SendResponse(ctx, fiber.StatusOK, products, "")
}

// GetProductByID Handler
func (c *ProductController) GetProductByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	product, err := c.service.GetProductByID(context.Background(), id)
	if err != nil {
		return utils.SendResponse(ctx, fiber.StatusNotFound, nil, "Product not Found")
	}
	return ctx.JSON(product)
}

func (c *ProductController) UpdateProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	exists, err := c.service.GetProductByID(context.Background(), id)
	if exists == nil || err != nil {
		return utils.SendResponse(ctx, fiber.StatusNotFound, nil, "Product not Found")
	}

	// Get validated DTO from the Local storage
	dto := ctx.Locals("validatedDTO")

	// Convert DTO to map for update
	updateData := utils.StructToMap(dto)

	// Update product in the database
	updatedProduct, err := c.service.UpdateProduct(context.Background(), id, bson.M(updateData))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{
		"data":    updatedProduct,
		"message": "Product updated successfully",
	})
}

// DeleteProduct Handler
func (c *ProductController) DeleteProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.service.DeleteProduct(context.Background(), id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
