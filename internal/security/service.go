package security

import (
	"errors"

	"github.com/google/uuid"
)

var (
	errBadCredentials    = errors.New("bad credentials")
	errUserAlreadyExists = errors.New("user already exists")
)

type DefaultSecurityService struct {
	UserRepository UserRepository
}

type UserRepository interface {
	FindByEmail(email string) (User, error)
	Save(user User) error
}

type User struct {
	ID       uuid.UUID
	Email    string
	Password string
}

func NewSecurityService(userRepository UserRepository) *DefaultSecurityService {
	return &DefaultSecurityService{userRepository}
}

func (s *DefaultSecurityService) Login(email string, password string) (string, error) {
	user, err := s.UserRepository.FindByEmail(email)
	if err != nil {
		return "", errBadCredentials
	}

	if user.Password != password {
		return "", errBadCredentials
	}

	return generateToken(user.ID, user.Email)
}

func (s *DefaultSecurityService) Register(email string, password string) error {
	user, err := s.UserRepository.FindByEmail(email)
	if err != nil {
		return err
	}

	if user.ID != uuid.Nil {
		return errUserAlreadyExists
	}

	user = User{
		ID:       uuid.New(),
		Email:    email,
		Password: password,
	}

	return s.UserRepository.Save(user)
}
