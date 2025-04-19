package repositories

import (
	"fmt"

	"github.com/task-manager/database"
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

	fmt.Print(result.RowsAffected)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("Task not found.")
	}

	return nil
}
