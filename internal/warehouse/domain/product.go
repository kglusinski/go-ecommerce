package domain

import (
	"errors"

	"github.com/google/uuid"
)

// ProductsRepository is the interface that wraps the basic product repository operations.
type ProductsRepository interface {
	FindAll() ([]*Product, error)
	FindByID(id uuid.UUID) (*Product, error)
	Save(product *Product) error
	Delete(id uuid.UUID) error
}

// Product represents a product in the warehouse.
type Product struct {
	id     uuid.UUID
	name   string
	price  float64
	amount float64
}

var (
	ErrInsufficientAmount       = errors.New("insufficient amount")
	ErrProductNotFound          = errors.New("product not found")
	ErrInvalidProductParameters = errors.New("invalid product parameters")
)

// NewProduct creates a new product with the given parameters.
func NewProduct(name string, price, amount float64) (*Product, error) {
	if name == "" || price <= 0 || amount < 0 {
		return nil, ErrInvalidProductParameters
	}

	return &Product{
		id:     uuid.New(),
		name:   name,
		price:  price,
		amount: amount,
	}, nil
}

func (p *Product) ID() uuid.UUID {
	return p.id
}

// ChangePrice changes the price of the product. Price cannot be negative.
func (p *Product) ChangePrice(newPrice float64) error {
	if newPrice <= 0 {
		return ErrInvalidProductParameters
	}

	p.price = newPrice

	return nil
}

// Supply adds the given amount to the product. Amount cannot be negative.
func (p *Product) Supply(toAdd float64) error {
	if toAdd <= 0 {
		return ErrInvalidProductParameters
	}

	p.amount += toAdd

	return nil
}

// Take removes the given amount from the product. Amount cannot be negative.
func (p *Product) Take(toTake float64) error {
	if toTake <= 0 {
		return ErrInvalidProductParameters
	}

	if p.amount < toTake {
		return ErrInsufficientAmount
	}

	p.amount -= toTake

	return nil
}
