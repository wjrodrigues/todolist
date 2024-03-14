package item

import (
	"testing"
	listDb "todolist/infra/database/pg/list"
	userDb "todolist/infra/database/pg/user"
	"todolist/internal/entity/item"
	listEntity "todolist/internal/entity/list"
	userEntity "todolist/internal/entity/user"
	uuid "todolist/pkg/entity"
	"todolist/pkg/test"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndDeleteItemWithSuccess(t *testing.T) {
	userDB := userDb.NewUserDB(test.Conn(t))
	user, _ := userEntity.NewUser("Pedro", "item_test@email.com", "123")

	listDB := listDb.NewListDB(test.Conn(t))
	list := listEntity.NewList("Pay bankslips", "any list", "in_progress", *user)

	itemDB := NewItemDB(test.Conn(t))
	item := item.NewItem("Title item", "Description item", "pending", list.ID)

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
	userDB := userDb.NewUserDB(test.Conn(t))
	user, _ := userEntity.NewUser("Pedro", "item_test@email.com", "123")

	listDB := listDb.NewListDB(test.Conn(t))
	list := listEntity.NewList("Pay bankslips", "any list", "in_progress", *user)

	itemDB := NewItemDB(test.Conn(t))
	item := item.NewItem("Title item", "Description item", "pending", list.ID)

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
	userDB := userDb.NewUserDB(test.Conn(t))
	user, _ := userEntity.NewUser("Pedro", "item_test@email.com", "123")

	listDB := listDb.NewListDB(test.Conn(t))
	list := listEntity.NewList("Pay bankslips", "any list", "in_progress", *user)

	itemDB := NewItemDB(test.Conn(t))
	item := item.NewItem("Title item", "Description item", "pending", list.ID)

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
	item := item.NewItem("Title item", "Description item", "pending", uuid.NewID())

	result, err := itemDB.FindById(item.ID)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}
