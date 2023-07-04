package cart

import (
	"context"
	"github.com/google/uuid"
)

type OrderService interface {
	MakeOrder(ctx context.Context, items []Item) (uuid.UUID, error)
}

type Item struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	Quantity int64     `json:"quantity"`
}

type Cart struct {
	ID    uuid.UUID `json:"id"`
	Items []Item    `json:"items"`

	order OrderService
}

func NewCart(os OrderService) *Cart {
	return &Cart{
		ID:    uuid.New(),
		Items: []Item{},
		order: os,
	}
}

func (c *Cart) AddItem(item Item) {
	c.Items = append(c.Items, item)
}

func (c *Cart) Finish(ctx context.Context) (uuid.UUID, error) {
	return c.order.MakeOrder(ctx, c.Items)
}
