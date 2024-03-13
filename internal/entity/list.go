package entity

import (
	"time"
	"todolist/pkg/entity"
)

type List struct {
	ID          entity.ID
	Title       string
	Description string
	Status      string
	Owner       User
	Items       []Item
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewList(title, description, status string, owner User) *List {
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

func (l *List) AddItem(item Item) {
	l.Items = append(l.Items, item)
}
