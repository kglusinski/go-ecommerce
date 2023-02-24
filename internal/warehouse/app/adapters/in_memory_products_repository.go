package adapters

import (
	"github.com/google/uuid"
	"github.com/inzkawka/go-ecommerce/internal/warehouse/domain"
)

type InMemoryProductsRepository struct {
	products map[uuid.UUID]*domain.Product
}

func NewInMemoryProductsRepository() *InMemoryProductsRepository {
	return &InMemoryProductsRepository{
		products: make(map[uuid.UUID]*domain.Product),
	}
}

func (i *InMemoryProductsRepository) FindAll() ([]*domain.Product, error) {
	var products []*domain.Product
	for _, product := range i.products {
		products = append(products, product)
	}

	return products, nil
}

func (i *InMemoryProductsRepository) FindByID(id uuid.UUID) (*domain.Product, error) {
	product, ok := i.products[id]
	if !ok {
		return nil, domain.ErrProductNotFound
	}

	return product, nil
}

func (i *InMemoryProductsRepository) Save(product *domain.Product) error {
	i.products[product.ID()] = product

	return nil
}

func (i *InMemoryProductsRepository) Delete(id uuid.UUID) error {
	delete(i.products, id)

	return nil
}
