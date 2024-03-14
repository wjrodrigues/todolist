package item

import (
	"errors"
	"testing"
	itemDb "todolist/infra/database/pg/item"
	listDb "todolist/infra/database/pg/list"
	"todolist/infra/database/pg/user"
	"todolist/internal/entity/item"
	"todolist/internal/entity/list"
	userEntity "todolist/internal/entity/user"
	"todolist/pkg/test"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndFindItemWithSuccess(t *testing.T) {
	userDB := user.NewUserDB(test.Conn(t))
	owner, _ := userEntity.NewUser("Pedro", "service_create_item@email.com", "123")
	userDB.Create(owner)

	list := list.NewList("Title list", "Description list", list.PENDING, *owner)
	listDB := listDb.NewListDB(test.Conn(t))
	listDB.Create(list)

	itemDB := itemDb.NewItemDB(test.Conn(t))
	itemInstance := item.NewItem("Item service", "Test item service", item.PENDING, list.ID)
	service := NewItemService(itemDB)

	err := service.Create(itemInstance)
	assert.Nil(t, err)

	result, err := service.Find(itemInstance.ID)
	assert.Nil(t, err)
	assert.Equal(t, result.Title, itemInstance.Title)

	t.Cleanup(func() {
		itemDB.Delete(itemInstance.ID)
		listDB.Delete(list.ID)
		userDB.Delete(owner.Email)
	})
}

func TestCreateItemWithFailed(t *testing.T) {
	owner, _ := userEntity.NewUser("Pedro", "service_create_item@email.com", "123")
	list := list.NewList("Title list", "Description list", list.PENDING, *owner)

	itemDB := itemDb.NewItemDB(test.Conn(t))
	itemInstance := item.NewItem("Item service", "Test item service", item.PENDING, list.ID)
	service := NewItemService(itemDB)

	err := service.Create(itemInstance)

	assert.Error(t, err)
}

func TestFindItemWithFailed(t *testing.T) {
	owner, _ := userEntity.NewUser("Pedro", "service_create_item@email.com", "123")
	list := list.NewList("Title list", "Description list", list.PENDING, *owner)

	itemDB := itemDb.NewItemDB(test.Conn(t))
	itemInstance := item.NewItem("Item service", "Test item service", item.PENDING, list.ID)
	service := NewItemService(itemDB)
	itemDB.DB.Close()

	_, err := service.Find(itemInstance.ID)

	assert.Error(t, err)
}

func TestUpdateStateItemWithSuccess(t *testing.T) {
	userDB := user.NewUserDB(test.Conn(t))
	owner, _ := userEntity.NewUser("Pedro", "service_create_item@email.com", "123")
	userDB.Create(owner)

	list := list.NewList("Title list", "Description list", list.PENDING, *owner)
	listDB := listDb.NewListDB(test.Conn(t))
	listDB.Create(list)

	itemDB := itemDb.NewItemDB(test.Conn(t))
	itemInstance := item.NewItem("Item service", "Test item service", item.PENDING, list.ID)
	service := NewItemService(itemDB)
	service.Create(itemInstance)

	service.ChangeStatus(*itemInstance, item.CANCELED)
	result, _ := service.Find(itemInstance.ID)
	assert.Equal(t, result.Status, item.CANCELED)

	service.ChangeStatus(*itemInstance, item.PENDING)
	result, _ = service.Find(itemInstance.ID)
	assert.Equal(t, result.Status, item.PENDING)

	service.ChangeStatus(*itemInstance, item.IN_PROGRESS)
	result, _ = service.Find(itemInstance.ID)
	assert.Equal(t, result.Status, item.IN_PROGRESS)

	service.ChangeStatus(*itemInstance, item.COMPLETED)
	result, _ = service.Find(itemInstance.ID)
	assert.Equal(t, result.Status, item.COMPLETED)

	t.Cleanup(func() {
		itemDB.Delete(itemInstance.ID)
		listDB.Delete(list.ID)
		userDB.Delete(owner.Email)
	})
}

func TestUpdateStateItemWithInvalidStatus(t *testing.T) {
	owner, _ := userEntity.NewUser("Pedro", "service_create_item@email.com", "123")

	list := list.NewList("Title list", "Description list", list.PENDING, *owner)

	itemDB := itemDb.NewItemDB(test.Conn(t))
	itemInstance := item.NewItem("Item service", "Test item service", item.PENDING, list.ID)
	service := NewItemService(itemDB)

	err := service.ChangeStatus(*itemInstance, "Any")

	assert.Equal(t, err, errors.New("no Any it's a valid status"))
}
