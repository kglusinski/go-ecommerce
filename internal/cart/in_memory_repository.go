package cart

import (
	"context"
	"github.com/google/uuid"
)

type CartRepository interface {
	New(ctx context.Context, cart *Cart)
	Get(ctx context.Context, id uuid.UUID) *Cart
	Update(ctx context.Context, cart *Cart)
	Delete(ctx context.Context, id uuid.UUID)
}

type InMemoryCartRepository struct {
	Carts map[uuid.UUID]*Cart
}

func NewInMemoryCartRepository() *InMemoryCartRepository {
	return &InMemoryCartRepository{
		Carts: map[uuid.UUID]*Cart{},
	}
}

func (r *InMemoryCartRepository) New(ctx context.Context, cart *Cart) {
	r.Carts[cart.ID] = cart
}

func (r *InMemoryCartRepository) Get(ctx context.Context, id uuid.UUID) *Cart {
	return r.Carts[id]
}

func (r *InMemoryCartRepository) Update(ctx context.Context, cart *Cart) {
	r.Carts[cart.ID] = cart
}

func (r *InMemoryCartRepository) Delete(ctx context.Context, id uuid.UUID) {
	delete(r.Carts, id)
}
