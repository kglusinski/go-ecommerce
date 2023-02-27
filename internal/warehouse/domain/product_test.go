package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	t.Run("should return error when name is empty", func(t *testing.T) {
		_, err := NewProduct("", 1, 1)
		assert.Error(t, err)
	})

	t.Run("should return error when price is negative", func(t *testing.T) {
		_, err := NewProduct("test", -1, 1)
		assert.Error(t, err)
	})

	t.Run("should return error when amount is negative", func(t *testing.T) {
		_, err := NewProduct("test", 1, -1)
		assert.Error(t, err)
	})

	t.Run("should return product when parameters are valid", func(t *testing.T) {
		product, err := NewProduct("test", 1, 1)
		assert.NoError(t, err)
		assert.Equal(t, "test", product.name)
		assert.Equal(t, 1.0, product.price)
		assert.Equal(t, 1.0, product.amount)
		assert.NotEmpty(t, product.id)
	})
}

func TestChangePrice(t *testing.T) {
	t.Run("should return error when price is negative", func(t *testing.T) {
		product, _ := NewProduct("test", 1, 1)
		err := product.ChangePrice(-1)
		assert.Error(t, err)
		assert.Equal(t, 1.0, product.price)
	})

	t.Run("should change product price when price is valid", func(t *testing.T) {
		product, _ := NewProduct("test", 1, 1)
		err := product.ChangePrice(2)
		assert.NoError(t, err)
		assert.Equal(t, 2.0, product.price)
	})
}

func TestSupply(t *testing.T) {
	t.Run("should return error when amount is negative", func(t *testing.T) {
		product, _ := NewProduct("test", 1, 1)
		err := product.Supply(-1)
		assert.Error(t, err)
		assert.Equal(t, 1.0, product.amount)
		assert.Equal(t, 1.0, product.price)
	})

	t.Run("should supply product when amount is valid", func(t *testing.T) {
		product, _ := NewProduct("test", 1, 1)
		err := product.Supply(2)
		assert.NoError(t, err)
		assert.Equal(t, 3.0, product.amount)
		assert.Equal(t, 1.0, product.price)
	})
}

func TestTake(t *testing.T) {
	t.Run("should return error when amount is negative", func(t *testing.T) {
		product, _ := NewProduct("test", 1, 1)
		err := product.Take(-1)
		assert.Error(t, err)
		assert.Equal(t, 1.0, product.amount)
		assert.Equal(t, 1.0, product.price)
	})

	t.Run("should return error when amount is greater than product amount", func(t *testing.T) {
		product, _ := NewProduct("test", 1, 1)
		err := product.Take(2)
		assert.Error(t, err)
		assert.Equal(t, 1.0, product.amount)
		assert.Equal(t, 1.0, product.price)
	})

	t.Run("should take product when amount is valid", func(t *testing.T) {
		product, _ := NewProduct("test", 1, 1)
		err := product.Take(1)
		assert.NoError(t, err)
		assert.Equal(t, 0.0, product.amount)
		assert.Equal(t, 1.0, product.price)
	})
}
