package domain

import "github.com/google/uuid"

type User struct {
	id              uuid.UUID
	email           string
	password        string
	deliveryAddress Address
}

type Address struct {
	firstLine  string
	secondLine string
	city       string
	postalCode string
	country    string
}

func NewUser(id uuid.UUID, email string, password string, deliveryAddress Address) *User {
	return &User{
		id:              id,
		email:           email,
		password:        password,
		deliveryAddress: deliveryAddress,
	}
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Email() string {
	return u.email
}
