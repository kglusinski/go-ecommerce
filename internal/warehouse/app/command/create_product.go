package command

import (
	"github.com/google/uuid"
	"github.com/inzkawka/go-ecommerce/internal/warehouse/domain"
)

type CreateProduct struct {
	Name   string
	Price  float64
	Amount float64
}

type CreateProductHandler struct {
	repo domain.ProductsRepository
}

func NewCreateProductHandler(repo domain.ProductsRepository) *CreateProductHandler {
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
