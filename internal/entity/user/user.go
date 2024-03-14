package user

import (
	"todolist/pkg/entity"
	"todolist/pkg/security"
)

type User struct {
	ID       entity.ID
	Name     string
	Email    string
	Password string
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := security.GenerateFromPassword(password)

	if err != nil {
		return &User{}, err
	}

	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}
