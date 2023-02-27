package command

import (
	"github.com/google/uuid"
	"github.com/inzkawka/go-ecommerce/internal/cart/domain"
)

type RemoveFromCart struct {
	CartID    uuid.UUID
	ProductID uuid.UUID
	Amount    float64
}

type RemoveFromCartHandler struct {
	repo domain.CartRepository
}

func NewRemoveFromCartHandler(repo domain.CartRepository) *RemoveFromCartHandler {
	return &RemoveFromCartHandler{repo: repo}
}

func (h *RemoveFromCartHandler) Handle(cmd RemoveFromCart) error {
	cart, err := h.repo.GetCart(cmd.CartID)
	if err != nil {
		return err
	}

	err = cart.RemoveItem(cmd.ProductID, cmd.Amount)
	if err != nil {
		return err
	}

	return h.repo.Save(cart)
}
