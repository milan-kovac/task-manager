package dtos

import "time"

type CreateTaskRequest struct {
	Title       string    `json:"title" validate:"required,min=3"`
	Description string    `json:"description" validate:"required"`
	DueDate     time.Time `json:"dueDate" validate:"required"`
}
