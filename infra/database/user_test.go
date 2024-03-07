package database

import (
	"testing"
	"todolist/internal/entity"
	"todolist/pkg/test"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndDeleteUserWithSuccess(t *testing.T) {
	userDB := NewUserDB(test.Conn(t))
	user, _ := entity.NewUser("Paulo", "paulo@gmail.com", "123456")

	err := userDB.Create(user)

	assert.Nil(t, err)

	t.Cleanup(func() {
		userDB.Delete(user.Email)
	})
}

func TestCreateWithFailed(t *testing.T) {
	userDB := NewUserDB(test.Conn(t))
	user, _ := entity.NewUser("Paulo", "paulo@gmail.com", "123456")

	userDB.Create(user)
	err := userDB.Create(user)

	assert.NotNil(t, err)

	t.Cleanup(func() {
		userDB.Delete(user.Email)
	})
}

func TestDeleteWithFailed(t *testing.T) {
	userDB := NewUserDB(test.Conn(t))
	userDB.DB.Close()

	err := userDB.Delete("paulo@gmail.com")

	assert.NotNil(t, err)
}

func TestFindByEmailOrIdWithSuccess(t *testing.T) {
	userDB := NewUserDB(test.Conn(t))
	user, _ := entity.NewUser("Maria", "maria@gmail.com", "123456")

	userDB.Create(user)

	identifiers := []struct {
		Id    string
		Email string
	}{
		{Id: user.ID.String(), Email: ""},
		{Id: "", Email: user.Email},
	}

	for _, identifier := range identifiers {
		result, err := userDB.FindByEmailOrId(identifier.Id, identifier.Email)

		assert.Nil(t, err)
		assert.NotNil(t, result.ID)
		assert.Empty(t, result.Password)
		assert.Equal(t, result.Email, user.Email)
		assert.Equal(t, result.Name, user.Name)
	}

	t.Cleanup(func() {
		userDB.Delete(user.Email)
	})
}

func TestFindByEmailOrIdNotFound(t *testing.T) {
	userDB := NewUserDB(test.Conn(t))
	user, _ := entity.NewUser("Maria", "maria@gmail.com", "123456")

	identifiers := []struct {
		Id    string
		Email string
	}{
		{Id: user.ID.String(), Email: ""},
		{Id: "", Email: user.Email},
	}

	for _, identifier := range identifiers {
		result, err := userDB.FindByEmailOrId(identifier.Id, identifier.Email)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	}
}
