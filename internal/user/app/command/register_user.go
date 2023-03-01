package command

import "github.com/inzkawka/go-ecommerce/internal/user/domain"

type RegisterUser struct {
	Email    string
	Password string
}

type RegisterUserHandler struct {
	repo domain.UserRepository
}

func NewRegisterUserHandler(repo domain.UserRepository) *RegisterUserHandler {
	return &RegisterUserHandler{
		repo: repo,
	}
}

func (h *RegisterUserHandler) Handle(cmd *RegisterUser) error {
	return nil
}
