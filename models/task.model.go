package models

import "time"

type Task struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Title       string
	Description string
	DueDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      uint `gorm:"index"`
}
