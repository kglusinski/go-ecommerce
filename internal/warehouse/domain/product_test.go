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
		assert.Equal(t, 1, product.price)
		assert.Equal(t, 1, product.amount)
		assert.NotEmpty(t, product.id)
	})
}
