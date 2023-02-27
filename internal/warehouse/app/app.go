package app

import (
	"errors"

	"github.com/inzkawka/go-ecommerce/internal/warehouse/app/command"
	"github.com/inzkawka/go-ecommerce/internal/warehouse/ports"
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
}

func NewApplication(repo ports.ProductsRepository) (*Application, error) {
	if repo == nil {
		return nil, ErrNilRepository
	}

	return &Application{
		Commands: Commands{
			CreateProduct: command.NewCreateProductHandler(repo),
		},
	}, nil
}
