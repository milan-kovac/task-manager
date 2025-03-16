package main

import (
	"github.com/task-manager/config"
	"github.com/task-manager/database"
)

func main() {
	config.ConfigureApp()

	db := &database.Database{}
	db.Connect()


	defer db.Close()
}