package app

import (
	"errors"

	"github.com/inzkawka/go-ecommerce/internal/user/app/command"
	"github.com/inzkawka/go-ecommerce/internal/user/domain"
)

var ErrNilRepository = errors.New("nil repository")

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	RegisterUser *command.RegisterUserHandler
}

type Queries struct {
}

func NewApplication(repo domain.UserRepository) (*Application, error) {
	if repo == nil {
		return nil, ErrNilRepository
	}

	return &Application{
		Commands: Commands{
			RegisterUser: command.NewRegisterUserHandler(repo),
		},
		Queries: Queries{},
	}, nil
}
