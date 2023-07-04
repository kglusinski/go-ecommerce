package order

import (
	"github.com/google/uuid"
)

type Product struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	Quantity int64     `json:"quantity"`
}

type Order struct {
	ID       uuid.UUID `json:"id"`
	Products []Product `json:"products"`
}

func NewOrder(products []Product) *Order {
	return &Order{
		ID:       uuid.New(),
		Products: products,
	}
}
