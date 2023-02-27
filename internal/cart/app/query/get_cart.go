package query

import (
	"github.com/google/uuid"
	"github.com/inzkawka/go-ecommerce/internal/cart/domain"
)

type GetCart struct {
	ID uuid.UUID
}

type GetCartHandler struct {
	repo domain.CartRepository
}

func NewGetCartHandler(repo domain.CartRepository) *GetCartHandler {
	return &GetCartHandler{repo: repo}
}

func (h *GetCartHandler) Handle(query *GetCart) (*domain.Cart, error) {
	return h.repo.GetCart(query.ID)
}
