package command

import (
	"github.com/google/uuid"
	"github.com/inzkawka/go-ecommerce/internal/cart/domain"
)

type CreateCart struct {
	UserID uuid.UUID
}

type CreateCartHandler struct {
	repo domain.CartRepository
}

func NewCreateCartHandler(repo domain.CartRepository) *CreateCartHandler {
	return &CreateCartHandler{repo: repo}
}

func (h *CreateCartHandler) Handle(cmd CreateCart) (uuid.UUID, error) {
	cart := domain.NewCart(cmd.UserID)

	return cart.ID(), h.repo.Save(cart)
}
