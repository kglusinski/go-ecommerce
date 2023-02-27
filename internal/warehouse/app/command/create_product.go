package command

import (
	"github.com/google/uuid"
	"github.com/inzkawka/go-ecommerce/internal/warehouse/domain"
	"github.com/inzkawka/go-ecommerce/internal/warehouse/ports"
)

type CreateProduct struct {
	Name   string
	Price  float64
	Amount float64
}

type CreateProductHandler struct {
	repo ports.ProductsRepository
}

func NewCreateProductHandler(repo ports.ProductsRepository) *CreateProductHandler {
	return &CreateProductHandler{
		repo: repo,
	}
}

func (h *CreateProductHandler) Handle(cmd CreateProduct) (uuid.UUID, error) {
	product, err := domain.NewProduct(cmd.Name, cmd.Price, cmd.Amount)
	if err != nil {
		return uuid.Nil, err
	}

	return product.ID(), h.repo.Save(product)
}
