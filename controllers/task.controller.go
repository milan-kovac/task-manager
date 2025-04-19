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

func GetTask(ctx *fiber.Ctx) error {
	id := ctx.Locals("id").(int)

	task, err := services.GetTask(id)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found.",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Task successfully fetched.",
		"task":    task,
	})
}

func GetTasks(ctx *fiber.Ctx) error {
	tasks, err := services.GetTasks()

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Tasks not found.",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Tasks successfully fetched.",
		"tasks":   tasks,
	})
}

func DeleteTask(ctx *fiber.Ctx) error {
	id := ctx.Locals("id").(int)

	err := services.DeleteTask(id)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found.",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Task successfully deleted.",
	})
}
