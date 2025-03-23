package models

import "time"

type User struct {
	ID        uint  `gorm:"primaryKey"`
	FullName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}