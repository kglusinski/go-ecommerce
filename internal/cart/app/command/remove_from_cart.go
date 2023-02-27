package command

import (
	"github.com/google/uuid"
	"github.com/inzkawka/go-ecommerce/internal/cart/domain"
)

type RemoveFromCart struct {
	CartID uuid.UUID
}

type RemoveFromCartHandler struct {
	repo domain.CartRepository
}

func NewRemoveFromCartHandler(repo domain.CartRepository) *RemoveFromCartHandler {
	return &RemoveFromCartHandler{repo: repo}
}

func (h *RemoveFromCartHandler) Handle(cmd RemoveFromCart) error {
	return nil
}
