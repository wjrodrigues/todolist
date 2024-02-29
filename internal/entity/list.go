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
	Items       []Item
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewList(title, description, status string, items []Item) *List {
	return &List{
		ID:          entity.NewID(),
		Title:       title,
		Description: description,
		Status:      status,
		Items:       items,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
