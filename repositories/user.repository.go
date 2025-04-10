package repositories

import (
	"github.com/task-manager/types"

	"github.com/task-manager/database"
	"github.com/task-manager/models"
)

func CreateUser(user *models.User) (*models.User, error) {
	if err := database.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserCredentialsByEmail(email string) (types.JWTUser, error) {
	var user models.User

	err := database.DB.Model(&models.User{}).
		Where("email = ?", email).
		Select("id, email, password").
		First(&user).Error

	if err != nil {
		return types.JWTUser{}, err
	}

	return types.JWTUser{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
