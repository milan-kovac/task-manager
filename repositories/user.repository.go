package repositories

import (
	"github.com/task-manager/database"
	"github.com/task-manager/models"
)


func  CreateUser(user *models.User) (*models.User, error) {
    if err := database.DB.Create(user).Error; err != nil {
        return nil, err
    }

    return user, nil
}