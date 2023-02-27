package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInsufficientAmount = errors.New("insufficient amount")
	ErrProductNotFound    = errors.New("product not found")
)

type Cart struct {
	id           uuid.UUID
	userID       uuid.UUID
	items        map[uuid.UUID]CartItem
	discountCode string
	totalPrice   float64
	finalPrice   float64
}

type CartItem struct {
	productID uuid.UUID
	amount    float64
	unitPrice float64
	sumPrice  float64
}

func NewCart(userID uuid.UUID) *Cart {
	return &Cart{
		id:     uuid.New(),
		userID: userID,
	}
}

func (c *Cart) ID() uuid.UUID {
	return c.id
}

func (c *Cart) UserID() uuid.UUID {
	return c.userID
}

func (c *Cart) Items() map[uuid.UUID]CartItem {
	return c.items
}

func (c *Cart) DiscountCode() string {
	return c.discountCode
}

func (c *Cart) TotalPrice() float64 {
	return c.totalPrice
}

func (c *Cart) FinalPrice() float64 {
	return c.finalPrice
}

func (c *Cart) AddItem(productID uuid.UUID, amount, unitPrice float64) error {
	if c.items == nil {
		c.items = make(map[uuid.UUID]CartItem, 0)
	}

	if amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	if unitPrice <= 0 {
		return errors.New("unit price must be greater than 0")
	}

	if item, ok := c.items[productID]; ok {
		item.amount += amount
		item.sumPrice += amount * unitPrice

		c.items[productID] = item
		c.totalPrice += amount * unitPrice

		return nil
	}

	c.items[productID] = CartItem{
		productID: productID,
		amount:    amount,
		unitPrice: unitPrice,
		sumPrice:  amount * unitPrice,
	}

	c.totalPrice += amount * unitPrice

	return nil
}

func (c *Cart) RemoveItem(productID uuid.UUID, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	item, ok := c.items[productID]
	if !ok {
		return ErrProductNotFound
	}

	if item.amount < amount {
		return ErrInsufficientAmount
	}

	if item.amount == amount {
		delete(c.items, productID)
		c.totalPrice -= amount * item.unitPrice

		return nil
	}

	item.amount -= amount
	item.sumPrice -= amount * item.unitPrice

	c.items[productID] = item
	c.totalPrice -= amount * item.unitPrice

	return nil
}
