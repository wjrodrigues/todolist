package database

import (
	"testing"
	"todolist/internal/entity"
	uuid "todolist/pkg/entity"
	"todolist/pkg/test"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndDeleteItemWithSuccess(t *testing.T) {
	itemDB := NewItemDB(test.Conn(t))
	item := entity.NewItem("Title item", "Description item", "pending")

	err := itemDB.Create(item)

	assert.Nil(t, err)

	t.Cleanup(func() {
		itemDB.Delete(item.ID)
	})
}

func TestCreateItemWithFailed(t *testing.T) {
	itemDB := NewItemDB(test.Conn(t))
	item := entity.NewItem("Title item", "Description item", "pending")

	itemDB.Create(item)
	err := itemDB.Create(item)

	assert.NotNil(t, err)

	t.Cleanup(func() {
		itemDB.Delete(item.ID)
	})
}

func TestDeletItemWithFailed(t *testing.T) {
	itemDB := NewItemDB(test.Conn(t))
	itemDB.DB.Close()

	err := itemDB.Delete(uuid.NewID())

	assert.NotNil(t, err)
}

func TestFindItemByIdWithSuccess(t *testing.T) {
	itemDB := NewItemDB(test.Conn(t))
	item := entity.NewItem("Title item", "Description item", "pending")

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

func TestFindItemByIdNotFound(t *testing.T) {
	itemDB := NewItemDB(test.Conn(t))
	item := entity.NewItem("Title item", "Description item", "pending")

	result, err := itemDB.FindById(item.ID)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}
