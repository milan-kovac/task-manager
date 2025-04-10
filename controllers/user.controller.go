package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/task-manager/dtos"
	"github.com/task-manager/services"
)

func Register(ctx *fiber.Ctx) error {
	var registerRequest dtos.RegisterRequest = ctx.Locals("body").(dtos.RegisterRequest)

	user, err := services.Register(registerRequest)
	if err != nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "User successfully registered.",
		"user":    user,
	})
}

func Login(ctx *fiber.Ctx) error {
	var loginRequest dtos.LoginRequest = ctx.Locals("body").(dtos.LoginRequest)

	token, err := services.Login(loginRequest)

	if err != nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "User successfully logged in.",
		"token":   token,
	})
}
