package ports

import (
	"github.com/google/uuid"
	"github.com/inzkawka/go-ecommerce/internal/warehouse/domain"
)

// ProductsRepository is the interface that wraps the basic product repository operations.
type ProductsRepository interface {
	FindAll() ([]*domain.Product, error)
	FindByID(id uuid.UUID) (*domain.Product, error)
	Save(product *domain.Product) error
	Delete(id uuid.UUID) error
}
