package domain

type UserRepository interface {
	FindByEmail(email string) (User, error)
	Save(user User) error
}
