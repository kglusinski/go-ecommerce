package domain

import "github.com/google/uuid"

type CartRepository interface {
	GetCart(id uuid.UUID) (*Cart, error)
	Save(cart *Cart) error
}
