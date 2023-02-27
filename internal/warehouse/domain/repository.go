package domain

import (
	"github.com/google/uuid"
)

// ProductsRepository is the interface that wraps the basic product repository operations.
type ProductsRepository interface {
	FindAll() ([]*Product, error)
	FindByID(id uuid.UUID) (*Product, error)
	Save(product *Product) error
	Delete(id uuid.UUID) error
}
