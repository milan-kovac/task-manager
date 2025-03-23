package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/task-manager/dtos"
	"github.com/task-manager/services"
)


func RegisterUser(ctx *fiber.Ctx) error {
	var registerRequest dtos.RegisterRequest = ctx.Locals("body").(dtos.RegisterRequest)

	user, err := services.RegisterUser(registerRequest)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

	return ctx.JSON(fiber.Map{
        "message": "User successfully registered",
        "user":    user,
    })
}