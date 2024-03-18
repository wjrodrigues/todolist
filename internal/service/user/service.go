package user

import (
	"todolist/infra/database"
	entity "todolist/internal/entity/user"
)

type IUserService interface {
	Create(user *entity.User) error
	Find(value string) (*entity.User, error)
}

type UserService struct {
	database database.IUser
}

func NewUserService(database database.IUser) *UserService {
	return &UserService{
		database: database,
	}
}

func (i *UserService) Create(list *entity.User) error {
	if err := i.database.Create(list); err != nil {
		return err
	}

	return nil
}

func (i *UserService) Find(value string) (*entity.User, error) {
	result, err := i.database.FindByEmailOrId(value, value)

	if err != nil {
		return nil, err
	}

	return result, nil
}
