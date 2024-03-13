package database

import (
	"testing"
	"todolist/internal/entity"
	uuid "todolist/pkg/entity"
	"todolist/pkg/test"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndDeleteItemWithSuccess(t *testing.T) {
	userDB := NewUserDB(test.Conn(t))
	user, _ := entity.NewUser("Pedro", "pedro@email.com", "123")

	listDB := NewListDB(test.Conn(t))
	list := entity.NewList("Pay bankslips", "any list", "in_progress", *user)

	itemDB := NewItemDB(test.Conn(t))
	item := entity.NewItem("Title item", "Description item", "pending", *list)

	userDB.Create(user)
	listDB.Create(list)
	err := itemDB.Create(item)

	assert.Nil(t, err)

	t.Cleanup(func() {
		itemDB.Delete(item.ID)
		listDB.Delete(list.ID)
		userDB.Delete(user.Email)
	})
}

func TestCreateItemWithFailed(t *testing.T) {
	userDB := NewUserDB(test.Conn(t))
	user, _ := entity.NewUser("Pedro", "pedro@email.com", "123")

	listDB := NewListDB(test.Conn(t))
	list := entity.NewList("Pay bankslips", "any list", "in_progress", *user)

	itemDB := NewItemDB(test.Conn(t))
	item := entity.NewItem("Title item", "Description item", "pending", *list)

	userDB.Create(user)
	listDB.Create(list)
	itemDB.Create(item)
	err := itemDB.Create(item)

	assert.NotNil(t, err)

	t.Cleanup(func() {
		itemDB.Delete(item.ID)
		listDB.Delete(list.ID)
		userDB.Delete(user.Email)
	})
}

func TestDeletItemWithFailed(t *testing.T) {
	itemDB := NewItemDB(test.Conn(t))
	itemDB.DB.Close()

	err := itemDB.Delete(uuid.NewID())

	assert.NotNil(t, err)
}

func TestFindItemByIdWithSuccess(t *testing.T) {
	userDB := NewUserDB(test.Conn(t))
	user, _ := entity.NewUser("Pedro", "pedro@email.com", "123")

	listDB := NewListDB(test.Conn(t))
	list := entity.NewList("Pay bankslips", "any list", "in_progress", *user)

	itemDB := NewItemDB(test.Conn(t))
	item := entity.NewItem("Title item", "Description item", "pending", *list)

	userDB.Create(user)
	listDB.Create(list)
	itemDB.Create(item)

	result, err := itemDB.FindById(item.ID)

	assert.Nil(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, result.Title, item.Title)
	assert.Equal(t, result.Description, item.Description)
	assert.Equal(t, result.Status, item.Status)

	t.Cleanup(func() {
		itemDB.Delete(item.ID)
		listDB.Delete(list.ID)
		userDB.Delete(user.Email)
	})
}

func TestFindItemByIdNotFound(t *testing.T) {
	itemDB := NewItemDB(test.Conn(t))
	item := entity.NewItem("Title item", "Description item", "pending", entity.List{})

	result, err := itemDB.FindById(item.ID)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}
