package adapters

import (
	"github.com/google/uuid"
	"github.com/inzkawka/go-ecommerce/internal/user/domain"
)

type InMemoryUserRepository struct {
	users map[uuid.UUID]*domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[uuid.UUID]*domain.User),
	}
}

func (r *InMemoryUserRepository) Save(user *domain.User) error {
	r.users[user.ID()] = user
	return nil
}

func (r *InMemoryUserRepository) FindByID(id uuid.UUID) (*domain.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, nil
	}
	return user, nil
}

func (r *InMemoryUserRepository) FindByEmail(email string) (*domain.User, error) {
	for _, user := range r.users {
		if user.Email() == email {
			return user, nil
		}
	}
	return nil, nil
}
