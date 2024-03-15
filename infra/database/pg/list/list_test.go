package list

import (
	"testing"
	userDb "todolist/infra/database/pg/user"
	"todolist/internal/entity/list"
	userEntity "todolist/internal/entity/user"
	uuid "todolist/pkg/entity"
	"todolist/pkg/test"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndDeleteListWithSuccess(t *testing.T) {
	listDB := NewListDB(test.Conn(t))
	userDB := userDb.NewUserDB(test.Conn(t))
	owner, _ := userEntity.NewUser("Pedro", "list_test@email.com", "123")
	list := list.NewList("Title list", "Description list", "pending", *owner)

	userDB.Create(owner)
	err := listDB.Create(list)

	assert.Nil(t, err)

	t.Cleanup(func() {
		listDB.Delete(list.ID)
		userDB.Delete(owner.Email)
	})
}

func TestCreateListWithFailed(t *testing.T) {
	listDB := NewListDB(test.Conn(t))
	userDB := userDb.NewUserDB(test.Conn(t))
	owner, _ := userEntity.NewUser("Pedro", "list_test@email.com", "123")
	list := list.NewList("Title list", "Description list", "pending", *owner)

	userDB.Create(owner)
	listDB.Create(list)

	listDB.Create(list)
	err := listDB.Create(list)

	assert.NotNil(t, err)

	t.Cleanup(func() {
		listDB.Delete(list.ID)
		userDB.Delete(owner.Email)
	})
}

func TestDeletListWithFailed(t *testing.T) {
	listDB := NewListDB(test.Conn(t))
	listDB.DB.Close()

	err := listDB.Delete(uuid.NewID())

	assert.NotNil(t, err)
}

func TestFindListByIdWithSuccess(t *testing.T) {
	listDB := NewListDB(test.Conn(t))
	userDB := userDb.NewUserDB(test.Conn(t))
	owner, _ := userEntity.NewUser("Pedro", "list_test@email.com", "123")
	list := list.NewList("Title list", "Description list", "pending", *owner)

	userDB.Create(owner)
	listDB.Create(list)

	result, err := listDB.FindById(list.ID)

	assert.Nil(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, result.Title, list.Title)
	assert.Equal(t, result.Description, list.Description)
	assert.Equal(t, result.Status, list.Status)

	t.Cleanup(func() {
		listDB.Delete(list.ID)
		userDB.Delete(owner.Email)
	})
}

func TestFindListByIdNotFound(t *testing.T) {
	listDB := NewListDB(test.Conn(t))
	owner, _ := userEntity.NewUser("Pedro", "list_test@email.com", "123")
	list := list.NewList("Title list", "Description list", "pending", *owner)

	result, err := listDB.FindById(list.ID)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestUpdateStatusListWithSuccess(t *testing.T) {
	userDB := userDb.NewUserDB(test.Conn(t))
	user, _ := userEntity.NewUser("Pedro", "list_update_test@email.com", "123")
	userDB.Create(user)

	listDB := NewListDB(test.Conn(t))
	listInstance := list.NewList("Pay bankslips", "any list", list.PENDING, *user)
	listDB.Create(listInstance)

	listInstance.Completed()
	err := listDB.UpdateStatus(*listInstance)
	assert.Nil(t, err)

	listInstance, _ = listDB.FindById(listInstance.ID)

	assert.Equal(t, listInstance.Status, list.COMPLETED)

	t.Cleanup(func() {
		listDB.Delete(listInstance.ID)
		userDB.Delete(user.Email)
	})
}

func TestUpdateStatusItemWithFailed(t *testing.T) {
	userDB := userDb.NewUserDB(test.Conn(t))
	user, _ := userEntity.NewUser("Pedro", "list_update_test@email.com", "123")
	userDB.Create(user)

	listDB := NewListDB(test.Conn(t))
	listInstance := list.NewList("Pay bankslips", "any list", list.PENDING, *user)
	listDB.Create(listInstance)

	listInstance.Status = "any"
	err := listDB.UpdateStatus(*listInstance)

	assert.Error(t, err)

	t.Cleanup(func() {
		listDB.Delete(listInstance.ID)
		userDB.Delete(user.Email)
	})
}
