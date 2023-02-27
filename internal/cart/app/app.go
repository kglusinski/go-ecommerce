package app

import (
	"errors"

	"github.com/inzkawka/go-ecommerce/internal/cart/app/command"
	"github.com/inzkawka/go-ecommerce/internal/cart/app/query"
	"github.com/inzkawka/go-ecommerce/internal/cart/domain"
)

var ErrNilRepository = errors.New("nil repository")

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateCart     *command.CreateCartHandler
	AddToCart      *command.AddToCartHandler
	RemoveFromCart *command.RemoveFromCartHandler
}

type Queries struct {
	GetCart *query.GetCartHandler
}

func NewApplication(repo domain.CartRepository) (*Application, error) {
	if repo == nil {
		return nil, ErrNilRepository
	}

	return &Application{
		Commands: Commands{
			CreateCart:     command.NewCreateCartHandler(repo),
			AddToCart:      command.NewAddToCartHandler(repo),
			RemoveFromCart: command.NewRemoveFromCartHandler(repo),
		},
		Queries: Queries{
			GetCart: query.NewGetCartHandler(repo),
		},
	}, nil
}
