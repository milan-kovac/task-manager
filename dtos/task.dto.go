package dtos

import "time"

type CreateTaskRequest struct {
	Title       string    `json:"title" validate:"required,min=3"`
	Description string    `json:"description" validate:"required"`
	DueDate     time.Time `json:"dueDate" validate:"required"`
}

type UpdateTaskRequest struct {
	Title       *string    `json:"title" validate:"omitempty,min=3"`
	Description *string    `json:"description" validate:"omitempty"`
	DueDate     *time.Time `json:"dueDate" validate:"omitempty"`
}
