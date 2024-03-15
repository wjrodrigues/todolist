package list

import (
	"time"
	"todolist/internal/entity/item"
	"todolist/internal/entity/user"
	"todolist/pkg/entity"
)

const (
	PENDING     = "pending"
	IN_PROGRESS = "in_progress"
	CANCELED    = "canceled"
	COMPLETED   = "completed"
)

type List struct {
	ID          entity.ID
	Title       string
	Description string
	Status      string
	Owner       user.User
	Items       []*item.Item
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
	l.Items = append(l.Items, &item)
}

func (l *List) Pending() {
	l.Status = PENDING
	l.updateStatusItems()
}

func (l *List) InProgress() {
	l.Status = IN_PROGRESS
	l.updateStatusItems()
}

func (l *List) Canceled() {
	l.Status = CANCELED
	l.updateStatusItems()
}

func (l *List) Completed() {
	l.Status = COMPLETED
	l.updateStatusItems()
}

func (l *List) updateStatusItems() {
	for _, item := range l.Items {
		switch l.Status {
		case PENDING:
			item.Pending()
		case IN_PROGRESS:
			item.InProgress()
		case COMPLETED:
			item.Completed()
		case CANCELED:
			item.Canceled()
		}
	}
}
