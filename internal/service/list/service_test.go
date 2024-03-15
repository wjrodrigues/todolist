package list

import (
	"errors"
	"testing"
	listDb "todolist/infra/database/pg/list"
	userDb "todolist/infra/database/pg/user"
	listEntity "todolist/internal/entity/list"
	userEntity "todolist/internal/entity/user"
	"todolist/pkg/test"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndFindListWithSuccess(t *testing.T) {
	userDB := userDb.NewUserDB(test.Conn(t))
	owner, _ := userEntity.NewUser("Pedro", "service_create_list@email.com", "123")
	userDB.Create(owner)

	listInstance := listEntity.NewList("Title list", "Description list", listEntity.PENDING, *owner)
	listDB := listDb.NewListDB(test.Conn(t))

	service := NewListService(listDB)

	err := service.Create(listInstance)
	assert.Nil(t, err)

	result, err := service.Find(listInstance.ID)
	assert.Nil(t, err)
	assert.Equal(t, result.Title, listInstance.Title)

	t.Cleanup(func() {
		listDB.Delete(listInstance.ID)
		userDB.Delete(owner.Email)
	})
}

func TestCreateListWithFailed(t *testing.T) {
	owner, _ := userEntity.NewUser("Pedro", "service_create_list@email.com", "123")
	listDB := listDb.NewListDB(test.Conn(t))
	listInstance := listEntity.NewList("Title list", "Description list", listEntity.PENDING, *owner)

	service := NewListService(listDB)

	err := service.Create(listInstance)

	assert.Error(t, err)
}

func TestFindListWithFailed(t *testing.T) {
	owner, _ := userEntity.NewUser("Pedro", "service_create_list@email.com", "123")
	listDB := listDb.NewListDB(test.Conn(t))
	listInstance := listEntity.NewList("Title list", "Description list", listEntity.PENDING, *owner)

	service := NewListService(listDB)
	listDB.DB.Close()

	_, err := service.Find(listInstance.ID)

	assert.Error(t, err)
}

func TestUpdateStateListWithSuccess(t *testing.T) {
	userDB := userDb.NewUserDB(test.Conn(t))
	owner, _ := userEntity.NewUser("Pedro", "service_create_list@email.com", "123")
	userDB.Create(owner)

	listInstance := listEntity.NewList("Title list", "Description list", listEntity.PENDING, *owner)
	listDB := listDb.NewListDB(test.Conn(t))
	listDB.Create(listInstance)

	service := NewListService(listDB)
	service.Create(listInstance)

	service.ChangeStatus(*listInstance, listEntity.CANCELED)
	result, _ := service.Find(listInstance.ID)
	assert.Equal(t, result.Status, listEntity.CANCELED)

	service.ChangeStatus(*listInstance, listEntity.PENDING)
	result, _ = service.Find(listInstance.ID)
	assert.Equal(t, result.Status, listEntity.PENDING)

	service.ChangeStatus(*listInstance, listEntity.IN_PROGRESS)
	result, _ = service.Find(listInstance.ID)
	assert.Equal(t, result.Status, listEntity.IN_PROGRESS)

	service.ChangeStatus(*listInstance, listEntity.COMPLETED)
	result, _ = service.Find(listInstance.ID)
	assert.Equal(t, result.Status, listEntity.COMPLETED)

	t.Cleanup(func() {
		listDB.Delete(listInstance.ID)
		userDB.Delete(owner.Email)
	})
}

func TestUpdateStateListWithInvalidStatus(t *testing.T) {
	owner, _ := userEntity.NewUser("Pedro", "service_create_list@email.com", "123")

	listInstance := listEntity.NewList("Title list", "Description list", listEntity.PENDING, *owner)

	listDB := listDb.NewListDB(test.Conn(t))
	service := NewListService(listDB)

	err := service.ChangeStatus(*listInstance, "Any")

	assert.Equal(t, err, errors.New("no Any it's a valid status"))
}
