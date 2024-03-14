package item

import (
	"fmt"
	"todolist/infra/database"
	entity "todolist/internal/entity/item"
	uuid "todolist/pkg/entity"
)

type IItemService interface {
	Create(item *entity.Item) error
	Find(id uuid.ID) (*entity.Item, error)
	ChangeStatus(item entity.Item, status string) error
}

type ItemService struct {
	database database.IItem
}

func NewItemService(database database.IItem) *ItemService {
	return &ItemService{
		database: database,
	}
}

func (i *ItemService) Create(item *entity.Item) error {
	if err := i.database.Create(item); err != nil {
		return err
	}

	return nil
}

func (i *ItemService) Find(id uuid.ID) (*entity.Item, error) {
	result, err := i.database.FindById(id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (i *ItemService) ChangeStatus(item entity.Item, status string) error {
	var err error

	switch status {
	case entity.PENDING:
		item.Pending()
	case entity.IN_PROGRESS:
		item.InProgress()
	case entity.COMPLETED:
		item.Completed()
	case entity.CANCELED:
		item.Canceled()
	default:
		err = fmt.Errorf("no %s it's a valid status", status)
	}

	if err != nil {
		return err
	}

	err = i.database.UpdateStatus(item)

	return err
}
