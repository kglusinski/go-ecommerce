package domain

import "github.com/google/uuid"

type Cart struct {
	id           uuid.UUID
	userID       uuid.UUID
	items        []CartItem
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

func (c *Cart) Items() []CartItem {
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
	c.items = append(c.items, CartItem{
		productID: productID,
		amount:    amount,
		unitPrice: unitPrice,
		sumPrice:  amount * unitPrice,
	})
	c.totalPrice += amount * unitPrice

	return nil
}

func (c *Cart) RemoveItem(productID uuid.UUID) error {
	for i, item := range c.items {
		if item.productID == productID {
			c.items = append(c.items[:i], c.items[i+1:]...)
			c.totalPrice -= item.sumPrice
			return nil
		}
	}

	return nil
}
