package services

import (
	"github.com/task-manager/dtos"
	"github.com/task-manager/models"
	"github.com/task-manager/repositories"
	"golang.org/x/crypto/bcrypt"
)


func  RegisterUser(registerRequest dtos.RegisterRequest) (*models.User, error) {
	hashedPassword, err := hashPassword(registerRequest.Password)
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


func hashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}