package services

import (
	"github.com/task-manager/dtos"
	"github.com/task-manager/models"
	"github.com/task-manager/repositories"
)

func CreateTask(createTaskRequest dtos.CreateTaskRequest, userId uint) (*models.Task, error) {
	task := &models.Task{
		Title:       createTaskRequest.Title,
		Description: createTaskRequest.Description,
		DueDate:     createTaskRequest.DueDate,
		UserID:      userId,
	}

	createdTask, err := repositories.CreateTask(task)

	if err != nil {
		return nil, err
	}

	return createdTask, nil
}
