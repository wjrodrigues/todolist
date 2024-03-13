package entity

import (
	"time"
	"todolist/pkg/entity"
)

type Item struct {
	ID          entity.ID
	Title       string
	Description string
	Status      string
	List        List
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewItem(title, description, status string, list List) *Item {
	return &Item{
		ID:          entity.NewID(),
		Title:       title,
		Description: description,
		Status:      status,
		List:        list,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
