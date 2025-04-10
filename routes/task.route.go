package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/task-manager/controllers"
	"github.com/task-manager/dtos"
	"github.com/task-manager/middlewares"
)

func RegisterTaskRoutes(app *fiber.App) {
	taskRoutes := app.Group("/tasks")

	taskRoutes.Post(
		"/",
		middlewares.AuthMiddleware(),
		middlewares.ValidateBody[dtos.CreateTaskRequest](),
		controllers.CreateTask)
}
