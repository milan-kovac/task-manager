package repositories

import (
	"github.com/task-manager/database"
	"github.com/task-manager/models"
)

func CreateTask(task *models.Task) (*models.Task, error) {
	if err := database.DB.Create(task).Error; err != nil {
		return nil, err
	}

	return task, nil
}
