package user

import (
	"testing"

	userDb "todolist/infra/database/pg/user"
	userEntity "todolist/internal/entity/user"
	"todolist/pkg/test"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndFindUsertWithSuccess(t *testing.T) {
	userDB := userDb.NewUserDB(test.Conn(t))
	user, _ := userEntity.NewUser("Pedro", "service_create_user@email.com", "123")

	service := NewUserService(userDB)

	err := service.Create(user)
	assert.Nil(t, err)

	result, err := service.Find(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, result.Name, user.Name)

	t.Cleanup(func() {
		userDB.Delete(user.Email)
	})
}

func TestCreateUsertWithFail(t *testing.T) {
	userDB := userDb.NewUserDB(test.Conn(t))
	user, _ := userEntity.NewUser("Pedro", "service_create_user@email.com", "123")

	service := NewUserService(userDB)

	service.Create(user)
	err := service.Create(user)
	assert.Error(t, err)

	t.Cleanup(func() {
		userDB.Delete(user.Email)
	})
}

func TestFidUsertWithFail(t *testing.T) {
	userDB := userDb.NewUserDB(test.Conn(t))
	userDB.DB.Close()
	service := NewUserService(userDB)

	response, err := service.Find("any")

	assert.Nil(t, response)
	assert.Error(t, err)
}

func TestValidatePassword(t *testing.T) {
	userDB := userDb.NewUserDB(test.Conn(t))
	user, _ := userEntity.NewUser("Pedro", "service_auth_user@email.com", "123")

	service := NewUserService(userDB)
	service.Create(user)

	id, err := service.Auth("service_auth_user@email.com", "123")
	assert.Nil(t, err)
	assert.Equal(t, id, user.ID.String())

	id, err = service.Auth("service_auth_user@email.com", "1234")
	assert.Error(t, err, "email or password are invalid")
	assert.Empty(t, id)

	id, err = service.Auth("service_auth_user@email.co", "123")
	assert.Error(t, err, "email or password are invalid")
	assert.Empty(t, id)

	t.Cleanup(func() {
		userDB.Delete(user.Email)
	})
}
