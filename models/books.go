package models

type Book struct {
   ID string `json:"id" validate:"required"`
   Title string `json:"title" validate:"required"`
   Author string  `json:"author" validate:"required"`
}