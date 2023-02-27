package adapters

import (
	"github.com/google/uuid"
	"github.com/inzkawka/go-ecommerce/internal/cart/domain"
)

type InMemoryCartRepository struct {
	carts map[uuid.UUID]*domain.Cart
}

func NewInMemoryCartRepository() *InMemoryCartRepository {
	return &InMemoryCartRepository{
		carts: make(map[uuid.UUID]*domain.Cart),
	}
}

func (i *InMemoryCartRepository) GetCart(id uuid.UUID) (*domain.Cart, error) {
	if cart, ok := i.carts[id]; ok {
		return cart, nil
	}

	return nil, domain.ErrCartNotFound
}

func (i *InMemoryCartRepository) Save(cart *domain.Cart) error {
	i.carts[cart.ID()] = cart

	return nil
}
