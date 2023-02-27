package command

import (
	"github.com/google/uuid"
	"github.com/inzkawka/go-ecommerce/internal/cart/domain"
)

type AddToCart struct {
	CartID    string
	ProductID uuid.UUID
	Amount    float64
	UnitPrice float64
}

type AddToCartHandler struct {
	repo domain.CartRepository
}

func NewAddToCartHandler(repo domain.CartRepository) *AddToCartHandler {
	return &AddToCartHandler{repo: repo}
}

func (h *AddToCartHandler) Handle(cmd AddToCart) error {
	return nil
}
