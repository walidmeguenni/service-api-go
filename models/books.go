package models

type Book struct {
   ID          string  `gorm:"primaryKey" json:"id"`
   Name        string  `json:"name" validate:"required"`
   Description string  `json:"description" validate:"required"`
   Price       float64 `json:"price" validate:"required,gt=0"`
}