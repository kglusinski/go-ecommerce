package command

import (
	"github.com/google/uuid"
	"github.com/inzkawka/go-ecommerce/internal/cart/domain"
)

type AddToCart struct {
	CartID    uuid.UUID `json:"cart_id"`
	ProductID uuid.UUID `json:"product_id"`
	Amount    float64   `json:"amount"`
	UnitPrice float64   `json:"unit_price"`
}

type AddToCartHandler struct {
	repo domain.CartRepository
}

func NewAddToCartHandler(repo domain.CartRepository) *AddToCartHandler {
	return &AddToCartHandler{repo: repo}
}

func (h *AddToCartHandler) Handle(cmd AddToCart) error {
	cart, err := h.repo.GetCart(cmd.CartID)
	if err != nil {
		return err
	}

	err = cart.AddItem(cmd.ProductID, cmd.Amount, cmd.UnitPrice)
	if err != nil {
		return err
	}

	return h.repo.Save(cart)
}
