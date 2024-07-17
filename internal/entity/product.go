package entity

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID          string
	Name        string
	Description *string
	Price       float64
	Categories  []string
	CreatedAt   time.Time
}

func NewProduct(name string, description *string, price float64, categories []string) *Product {
	return &Product{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
		Price:       price,
		Categories:  categories,
		CreatedAt:   time.Now(),
	}
}
