package list

import (
	"time"
	"todolist/internal/entity/item"
	"todolist/internal/entity/user"
	"todolist/pkg/entity"
)

type List struct {
	ID          entity.ID
	Title       string
	Description string
	Status      string
	Owner       user.User
	Items       []item.Item
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewList(title, description, status string, owner user.User) *List {
	return &List{
		ID:          entity.NewID(),
		Title:       title,
		Description: description,
		Status:      status,
		Owner:       owner,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (l *List) AddItem(item item.Item) {
	l.Items = append(l.Items, item)
}
