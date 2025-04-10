package services

import (
	"errors"

	"github.com/task-manager/dtos"
	"github.com/task-manager/models"
	"github.com/task-manager/repositories"
)

func Register(registerRequest dtos.RegisterRequest) (*models.User, error) {
	hashedPassword, err := HashPassword(registerRequest.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		FullName: registerRequest.FullName,
		Email:    registerRequest.Email,
		Password: hashedPassword,
	}

	createdUser, err := repositories.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func Login(loginRequest dtos.LoginRequest) (string, error) {
	user, err := repositories.GetUserCredentialsByEmail(loginRequest.Email)
	if err != nil {
		return "", errors.New("Invalid email or password.")
	}

	if err := CheckPassword(loginRequest.Password, user.Password); err != nil {
		return "", err
	}

	token, err := GenerateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
