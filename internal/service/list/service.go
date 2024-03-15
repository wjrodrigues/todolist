package list

import (
	"fmt"
	"todolist/infra/database"
	entity "todolist/internal/entity/list"
	uuid "todolist/pkg/entity"
)

type IListService interface {
	Create(list *entity.List) error
	Find(id uuid.ID) (*entity.List, error)
	ChangeStatus(list entity.List, status string) error
}

type ListService struct {
	database database.IList
}

func NewListService(database database.IList) *ListService {
	return &ListService{
		database: database,
	}
}

func (i *ListService) Create(list *entity.List) error {
	if err := i.database.Create(list); err != nil {
		return err
	}

	return nil
}

func (i *ListService) Find(id uuid.ID) (*entity.List, error) {
	result, err := i.database.FindById(id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (i *ListService) ChangeStatus(list entity.List, status string) error {
	var err error

	switch status {
	case entity.PENDING:
		list.Pending()
	case entity.IN_PROGRESS:
		list.InProgress()
	case entity.COMPLETED:
		list.Completed()
	case entity.CANCELED:
		list.Canceled()
	default:
		err = fmt.Errorf("no %s it's a valid status", status)
	}

	if err != nil {
		return err
	}

	err = i.database.UpdateStatus(list)

	return err
}
