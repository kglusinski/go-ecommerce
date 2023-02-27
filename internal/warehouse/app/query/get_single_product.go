package query

import (
	"github.com/google/uuid"
	"github.com/inzkawka/go-ecommerce/internal/warehouse/domain"
)

type GetSingleProduct struct {
	ID uuid.UUID
}

type GetSingleProductHandler struct {
	repo domain.ProductsRepository
}

func NewGetSingleProductHandler(repo domain.ProductsRepository) *GetSingleProductHandler {
	return &GetSingleProductHandler{
		repo: repo,
	}
}

func (h *GetSingleProductHandler) Handle(cmd GetSingleProduct) (*domain.Product, error) {
	return h.repo.FindByID(cmd.ID)
}
