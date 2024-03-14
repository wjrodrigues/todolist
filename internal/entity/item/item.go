package item

import (
	"time"
	"todolist/pkg/entity"
)

const (
	PENDING     = "pending"
	IN_PROGRESS = "in_progress"
	CANCELED    = "canceled"
	COMPLETED   = "completed"
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

func (i *Item) Pending() {
	i.Status = PENDING
}

func (i *Item) InProgress() {
	i.Status = IN_PROGRESS
}

func (i *Item) Canceled() {
	i.Status = CANCELED
}

func (i *Item) Completed() {
	i.Status = COMPLETED
}
