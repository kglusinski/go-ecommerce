package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddItem(t *testing.T) {
	t.Run("should add item to cart", func(t *testing.T) {
		cart := NewCart(uuid.New())
		err := cart.AddItem(uuid.New(), 1, 1.0)

		assert.NoError(t, err)
	})

	t.Run("should add item to cart with existing product", func(t *testing.T) {
		cart := NewCart(uuid.New())
		productID := uuid.New()
		err := cart.AddItem(productID, 1, 1.0)
		assert.NoError(t, err)

		err = cart.AddItem(productID, 1, 1.0)
		assert.NoError(t, err)
		assert.Equal(t, 2.0, cart.items[productID].amount)
		assert.Equal(t, 2.0, cart.items[productID].sumPrice)
		assert.Equal(t, 2.0, cart.totalPrice)
	})

	t.Run("should add item to cart with existing other product", func(t *testing.T) {
		cart := NewCart(uuid.New())
		productID := uuid.New()
		err := cart.AddItem(productID, 1, 1.0)
		assert.NoError(t, err)

		err = cart.AddItem(uuid.New(), 1, 1.0)
		assert.NoError(t, err)
		assert.Equal(t, 2.0, cart.totalPrice)
	})

	t.Run("should return an error when adding item with zero amount", func(t *testing.T) {
		cart := NewCart(uuid.New())
		err := cart.AddItem(uuid.New(), 0, 1.0)

		assert.Error(t, err)
	})

	t.Run("should return an error when adding item with zero unit price", func(t *testing.T) {
		cart := NewCart(uuid.New())
		err := cart.AddItem(uuid.New(), 1, 0.0)

		assert.Error(t, err)
	})

	t.Run("should return an error when adding item with negative amount", func(t *testing.T) {
		cart := NewCart(uuid.New())
		err := cart.AddItem(uuid.New(), -1, 1.0)

		assert.Error(t, err)
	})

	t.Run("should return an error when adding item with negative unit price", func(t *testing.T) {
		cart := NewCart(uuid.New())
		err := cart.AddItem(uuid.New(), 1, -1.0)

		assert.Error(t, err)
	})
}

func TestRemoveItem(t *testing.T) {
	t.Run("should remove item from cart", func(t *testing.T) {
		cart := NewCart(uuid.New())
		productID := uuid.New()
		err := cart.AddItem(productID, 1, 1.0)
		assert.NoError(t, err)

		err = cart.RemoveItem(productID, 1)
		assert.NoError(t, err)
		assert.Equal(t, 0.0, cart.totalPrice)
	})

	t.Run("should remove item from cart with existing other product", func(t *testing.T) {
		cart := NewCart(uuid.New())
		productID := uuid.New()
		err := cart.AddItem(productID, 1, 1.0)
		assert.NoError(t, err)
		err = cart.AddItem(uuid.New(), 1, 2.0)
		assert.NoError(t, err)

		err = cart.RemoveItem(productID, 1)
		assert.NoError(t, err)
		assert.Equal(t, 2.0, cart.totalPrice)
	})

	t.Run("should remove partially item from existing product", func(t *testing.T) {
		cart := NewCart(uuid.New())
		productID := uuid.New()
		err := cart.AddItem(productID, 2, 1.0)
		assert.NoError(t, err)

		err = cart.RemoveItem(productID, 1)
		assert.NoError(t, err)
		assert.Equal(t, 1.0, cart.items[productID].amount)
		assert.Equal(t, 1.0, cart.items[productID].sumPrice)
		assert.Equal(t, 1.0, cart.totalPrice)
	})

	t.Run("should remove all items from existing product", func(t *testing.T) {
		cart := NewCart(uuid.New())
		productID := uuid.New()
		err := cart.AddItem(productID, 2, 1.0)
		assert.NoError(t, err)

		err = cart.RemoveItem(productID, 2)
		assert.NoError(t, err)
		assert.Equal(t, 0.0, cart.items[productID].amount)
		assert.Equal(t, 0.0, cart.items[productID].sumPrice)
		assert.Equal(t, 0.0, cart.totalPrice)
	})

	t.Run("should remove all items from existing product when there is other product", func(t *testing.T) {
		cart := NewCart(uuid.New())
		productID := uuid.New()
		err := cart.AddItem(productID, 2, 1.0)
		assert.NoError(t, err)

		err = cart.AddItem(uuid.New(), 1, 1.0)
		assert.NoError(t, err)

		err = cart.RemoveItem(productID, 2)
		assert.NoError(t, err)
		assert.Equal(t, 0.0, cart.items[productID].amount)
		assert.Equal(t, 0.0, cart.items[productID].sumPrice)
		assert.Equal(t, 1.0, cart.totalPrice)
	})

	t.Run("should return an error when item not in the cart", func(t *testing.T) {
		cart := NewCart(uuid.New())
		err := cart.RemoveItem(uuid.New(), 1)

		assert.Error(t, err)
	})

	t.Run("should return an error when removing item with zero amount", func(t *testing.T) {
		cart := NewCart(uuid.New())
		productID := uuid.New()
		cart.AddItem(productID, 1, 1.0)
		err := cart.RemoveItem(productID, 0)

		assert.Error(t, err)
	})

	t.Run("should return an error when removing item with negative amount", func(t *testing.T) {
		cart := NewCart(uuid.New())
		productID := uuid.New()
		cart.AddItem(productID, 1, 1.0)
		err := cart.RemoveItem(productID, -1)

		assert.Error(t, err)
	})

	t.Run("should return an error when removing more items than in the cart", func(t *testing.T) {
		cart := NewCart(uuid.New())
		productID := uuid.New()
		cart.AddItem(productID, 1, 1.0)
		err := cart.RemoveItem(productID, 2)

		assert.Error(t, err)
	})
}
