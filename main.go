package main

import (
	"github.com/gofiber/fiber"
	"github.com/task-manager/config"
	"github.com/task-manager/database"
)

func main() {
	app := fiber.New()
	config.ConfigureApp()

	db := &database.Database{}
	db.Initialize()

	app.Listen(":3000")
}