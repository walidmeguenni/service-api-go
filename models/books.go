package models

import (
	"time"
)

type Book struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Price       float64    `json:"price" validate:"required,gt=0"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
