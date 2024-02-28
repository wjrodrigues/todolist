package entity

import (
	"todolist/pkg/entity"
	"todolist/pkg/security"
)

type UserInstance struct {
	ID       entity.ID
	Name     string
	Email    string
	Password string
}

func NewUser(name, email, password string) (*UserInstance, error) {
	hash, err := security.GenerateFromPassword(password)

	if err != nil {
		return &UserInstance{}, err
	}

	return &UserInstance{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}
