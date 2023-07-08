package controllers

import (
	"musixer/api/internal/app/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Get User Profile
// @Description Get the profile of the currently authenticated user
// @Tags Users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.UserResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 502 {object} models.ErrorResponse
// @Router /api/users/me [get]
func GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserResponse)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}
