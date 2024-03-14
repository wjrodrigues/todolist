package item

import (
	"time"
	"todolist/pkg/entity"
)

type Item struct {
	ID          entity.ID
	Title       string
	Description string
	Status      string
	ListId      entity.ID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewItem(title, description, status string, listId entity.ID) *Item {
	return &Item{
		ID:          entity.NewID(),
		Title:       title,
		Description: description,
		Status:      status,
		ListId:      listId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
