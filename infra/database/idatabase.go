package database

import (
	"todolist/internal/entity/item"
	"todolist/internal/entity/list"
	"todolist/internal/entity/user"

	uuid "todolist/pkg/entity"
)

type IUser interface {
	Create(user *user.User) error
	FindByEmailOrId(id, email string) (*user.User, error)
}

type IItem interface {
	Create(item *item.Item) error
	Delete(id uuid.ID) error
	FindById(id uuid.ID) (*item.Item, error)
	UpdateStatus(item item.Item) error
}

type IList interface {
	Create(list *list.List) error
	Delete(id uuid.ID) error
	FindById(id uuid.ID) (*list.List, error)
	UpdateStatus(list list.List) error
}
