package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/task-manager/controllers"
	"github.com/task-manager/dtos"
	"github.com/task-manager/middlewares"
)

func RegisterUserRoutes(app *fiber.App) {
	userRoutes := app.Group("/users")

	userRoutes.Post(
		"/register",
		 middlewares.ValidateBody[dtos.RegisterRequest](), 
		 controllers.RegisterUser)
}