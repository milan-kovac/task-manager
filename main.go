package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/task-manager/config"
	"github.com/task-manager/database"
	"github.com/task-manager/routes"
)

func main() {
	app := fiber.New()

	// APP CONFIG
	config.ConfigureApp()

	// DATABASE CONFIG
	database.Connect()

	// ROUTES
	routes.RegisterUserRoutes(app)

	app.Listen(":3000")
}