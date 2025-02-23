package controllers

import (
	"rest/services"
	"rest/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"

	dtos "rest/routes/dtos/auth"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{service}
}

// Register API
func (ctrl *AuthController) Register(c *fiber.Ctx) error {
	dto, _ := c.Locals("validatedDTO").(*dtos.RegisterDTO)

	token, err := ctrl.service.Register(dto)
	if err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, nil, err.Error())
	}

	return utils.SendResponse(c, fiber.StatusCreated, fiber.Map{"token": token}, "User registered successfully")
}

// Login API
func (ctrl *AuthController) Login(c *fiber.Ctx) error {
	dto, _ := c.Locals("validatedDTO").(*dtos.LoginDTO)

	token, err := ctrl.service.Login(dto.Phone, dto.Password)
	if err != nil {
		return utils.SendResponse(c, fiber.StatusUnauthorized, nil, err.Error())
	}

	return utils.SendResponse(c, fiber.StatusOK, fiber.Map{"token": token}, "")
}

// Update Profile API
func (ctrl *AuthController) UpdateProfile(c *fiber.Ctx) error {
	userID := c.Locals("userId").(string)
	dto, _ := c.Locals("validatedDTO").(*dtos.UpdateProfileDTO)

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, nil, err.Error())
	}

	if err := ctrl.service.UpdateProfile(objectID, dto); err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, nil, "failed to update profile")
	}

	return utils.SendResponse(c, fiber.StatusOK, dto, "profile updated")
}
