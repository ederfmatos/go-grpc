package entity

import (
	"github.com/google/uuid"
	"time"
)

type Category struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
}

func NewCategory(name, description string) *Category {
	return &Category{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
	}
}
