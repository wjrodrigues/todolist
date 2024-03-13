package database

import (
	"testing"
	"todolist/internal/entity"
	uuid "todolist/pkg/entity"
	"todolist/pkg/test"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndDeleteListWithSuccess(t *testing.T) {
	listDB := NewListDB(test.Conn(t))
	list := entity.NewList("Title list", "Description list", "pending", []entity.Item{})

	err := listDB.Create(list)

	assert.Nil(t, err)

	t.Cleanup(func() {
		listDB.Delete(list.ID)
	})
}

func TestCreateListWithFailed(t *testing.T) {
	listDB := NewListDB(test.Conn(t))
	list := entity.NewList("Title list", "Description list", "pending", []entity.Item{})

	listDB.Create(list)
	err := listDB.Create(list)

	assert.NotNil(t, err)

	t.Cleanup(func() {
		listDB.Delete(list.ID)
	})
}

func TestDeletListWithFailed(t *testing.T) {
	listDB := NewListDB(test.Conn(t))
	listDB.DB.Close()

	err := listDB.Delete(uuid.NewID())

	assert.NotNil(t, err)
}

func TestFindListByIdWithSuccess(t *testing.T) {
	itemDB := NewListDB(test.Conn(t))
	item := entity.NewList("Title list", "Description list", "pending", []entity.Item{})

	itemDB.Create(item)

	result, err := itemDB.FindById(item.ID)

	assert.Nil(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, result.Title, item.Title)
	assert.Equal(t, result.Description, item.Description)
	assert.Equal(t, result.Status, item.Status)

	t.Cleanup(func() {
		itemDB.Delete(item.ID)
	})
}

func TestFindListByIdNotFound(t *testing.T) {
	listDB := NewListDB(test.Conn(t))
	list := entity.NewList("Title list", "Description list", "pending", []entity.Item{})

	result, err := listDB.FindById(list.ID)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}
