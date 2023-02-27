package app

import (
	"errors"

	"github.com/inzkawka/go-ecommerce/internal/warehouse/app/command"
	"github.com/inzkawka/go-ecommerce/internal/warehouse/app/query"
	"github.com/inzkawka/go-ecommerce/internal/warehouse/domain"
)

var ErrNilRepository = errors.New("nil repository")

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateProduct *command.CreateProductHandler
}

type Queries struct {
	GetProduct *query.GetSingleProductHandler
}

func NewApplication(repo domain.ProductsRepository) (*Application, error) {
	if repo == nil {
		return nil, ErrNilRepository
	}

	return &Application{
		Commands: Commands{
			CreateProduct: command.NewCreateProductHandler(repo),
		},
		Queries: Queries{
			GetProduct: query.NewGetSingleProductHandler(repo),
		},
	}, nil
}
