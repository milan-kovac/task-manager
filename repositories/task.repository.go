package repositories

import (
	"fmt"

	"github.com/task-manager/database"
	"github.com/task-manager/dtos"
	"github.com/task-manager/models"
)

func CreateTask(task *models.Task) (*models.Task, error) {
	if err := database.DB.Create(task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func GetTaskById(id int) (*models.Task, error) {
	var task models.Task

	if err := database.DB.Where("id = ?", id).First(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task

	if err := database.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func DeleteTaskById(id int) error {
	result := database.DB.Delete(&models.Task{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("Task not found.")
	}

	return nil
}

func UpdateTask(id int, updateTaskRequest dtos.UpdateTaskRequest) (*models.Task, error) {
	if err := database.DB.Model(&models.Task{}).Where("id = ?", id).Updates(updateTaskRequest).Error; err != nil {
		return nil, err
	}

	var updatedTask models.Task
	if err := database.DB.First(&updatedTask, id).Error; err != nil {
		return nil, err
	}

	return &updatedTask, nil
}
