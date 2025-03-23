package dtos

type RegisterRequest struct {
	FullName string `json:"fullName" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}