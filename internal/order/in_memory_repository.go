package order

import "github.com/google/uuid"

type InMemoryOrderRepository struct {
	Orders map[uuid.UUID]*Order
}

func NewInMemoryOrderRepository() *InMemoryOrderRepository {
	return &InMemoryOrderRepository{
		Orders: map[uuid.UUID]*Order{},
	}
}

func (r *InMemoryOrderRepository) New(order *Order) {
	r.Orders[order.ID] = order
}

func (r *InMemoryOrderRepository) Get(id uuid.UUID) *Order {
	return r.Orders[id]
}

func (r *InMemoryOrderRepository) Update(order *Order) {
	r.Orders[order.ID] = order
}

func (r *InMemoryOrderRepository) Delete(id uuid.UUID) {
	delete(r.Orders, id)
}
