package database

import (
	"todolist/internal/entity/item"
	"todolist/internal/entity/list"
	"todolist/internal/entity/user"

	uuid "todolist/pkg/entity"
)

type User interface {
	Create(user *user.User) error
	Delete(email string) error
	FindByEmailOrId(id, email string) (*user.User, error)
}

type Item interface {
	Create(item *item.Item) error
	Delete(id uuid.ID) error
	FindById(id uuid.ID) (*item.Item, error)
}

type List interface {
	Create(list *list.List) error
	Delete(id uuid.ID) error
	FindById(id uuid.ID) (*list.List, error)
}
