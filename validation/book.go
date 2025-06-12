package validation

type CreateBook struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Price       float64    `json:"price" validate:"required,gt=0"`
}