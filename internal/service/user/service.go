package user

import (
	"errors"
	"todolist/infra/database"
	entity "todolist/internal/entity/user"
)

const InvalidAuth = "email or password are invalid"

type IUserService interface {
	Create(user *entity.User) error
	Find(value string) (*entity.User, error)
	Auth(email, password string) (string, error)
}

type UserService struct {
	database database.IUser
}

func NewUserService(database database.IUser) *UserService {
	return &UserService{
		database: database,
	}
}

func (u *UserService) Create(user *entity.User) error {
	if err := u.database.Create(user); err != nil {
		return err
	}

	return nil
}

func (u *UserService) Find(value string) (*entity.User, error) {
	result, err := u.database.FindByEmailOrId(value, value)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *UserService) Auth(email, password string) (string, error) {
	result, _ := u.database.FindByEmailOrId("", email)

	if result == nil {
		return "", errors.New(InvalidAuth)
	}

	valid := result.ValidatePassword(password)

	if !valid {
		return "", errors.New(InvalidAuth)
	}

	return result.ID.String(), nil
}
