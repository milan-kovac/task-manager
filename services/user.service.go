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

	createdUser, err := repositories.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func Login(loginRequest dtos.LoginRequest) (string, error) {
	password, err := repositories.GetPasswordByEmail(loginRequest.Email)
	if err != nil {
		return "", errors.New("Invalid email or password.")
	}

	if err := CheckPassword(loginRequest.Password, password); err != nil {
		return "", err
	}

	token, err := GenerateJWT(loginRequest.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
