package repositories

import (
	"github.com/task-manager/database"
	"github.com/task-manager/models"
)


func  Create(user *models.User) (*models.User, error) {
    if err := database.DB.Create(user).Error; err != nil {
        return nil, err
    }

    return user, nil
}


func GetPasswordByEmail(email string) (string, error) {
    var password string
    err := database.DB.Model(&models.User{}).Where("email = ?", email).Select("password").Scan(&password).Error
    if err != nil {
        return "", err
    }

    return password, nil
}

