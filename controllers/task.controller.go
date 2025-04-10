package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/task-manager/dtos"
	"github.com/task-manager/services"
	"github.com/task-manager/types"
)

func CreateTask(ctx *fiber.Ctx) error {
	var createTaskRequest dtos.CreateTaskRequest = ctx.Locals("body").(dtos.CreateTaskRequest)
	var currentUser types.CurrentUser = ctx.Locals("currentUser").(types.CurrentUser)

	task, err := services.CreateTask(createTaskRequest, currentUser.ID)
	if err != nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Task successfully created.",
		"task":    task,
	})
}
