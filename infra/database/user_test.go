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

// func TestFindByEmail(t *testing.T) {
// 	userDB := NewUser(test.Conn(t))
// 	user, _ := entity.NewUser("Maria", "maria@gmail.com", "123456")

// 	userDB.Create(user)
// 	result, err := userDB.FindByEmail(user.Email)

// 	assert.Nil(t, err)
// 	assert.NotNil(t, result.ID)
// 	assert.Empty(t, result.Password)
// 	assert.Equal(t, result.Email, user.Email)
// 	assert.Equal(t, result.Name, user.Name)

// 	t.Cleanup(func() {
// 		userDB.Delete(user.Email)
// 	})
// }

// func TestFindByID(t *testing.T) {
// 	userDB := NewUser(test.Conn(t))
// 	user, _ := entity.NewUser("Maria", "maria@gmail.com", "123456")

// 	userDB.Create(user)
// 	result, err := userDB.FindByID(user.ID.String())

// 	assert.Nil(t, err)
// 	assert.NotNil(t, result.ID)
// 	assert.Empty(t, result.Password)
// 	assert.Equal(t, result.Email, user.Email)
// 	assert.Equal(t, result.Name, user.Name)

// 	t.Cleanup(func() {
// 		userDB.Delete(user.Email)
// 	})
// }

// func TestNotFoundWhenFindByID(t *testing.T) {
// 	userDB := NewUser(test.Conn(t))
// 	result, err := userDB.FindByID("52fb6d91-29cb-43bb-965b-9dd37d482623")

// 	assert.Error(t, err)
// 	assert.Empty(t, result)
// }

// func TestCreateFailed(t *testing.T) {
// 	userDB := NewUser(test.Conn(t))
// 	user, _ := entity.NewUser("Paulo", "paulo@gmail.com", "123456")

// 	userDB.Create(user)
// 	err := userDB.Create(user)

// 	assert.Error(t, err)

// 	t.Cleanup(func() {
// 		userDB.Delete(user.Email)
// 	})
// }

// func TestDeleteFailed(t *testing.T) {
// 	db := test.Conn(t)
// 	userDB := NewUser(db)
// 	db.Close()

// 	err := userDB.Delete("")
// 	assert.Error(t, err)
// }

// func TestNotFoundWhenFindByEmail(t *testing.T) {
// 	userDB := NewUser(test.Conn(t))
// 	result, err := userDB.FindByEmail("carlos@gmail.com")

// 	assert.Error(t, err)
// 	assert.Empty(t, result)
// }

// func TestUpdateNameUser(t *testing.T) {
// 	userDB := NewUser(test.Conn(t))
// 	user, _ := entity.NewUser("Paulo", "paulo@gmail.com", "123456")
// 	userDB.Create(user)
// 	userValues, _ := entity.NewUser("Pedro", "", "")
// 	userValues.Password = ""

// 	err := userDB.Update(userValues, user.ID.String())
// 	userUpdated, _ := userDB.FindByID(user.ID.String())

// 	assert.Nil(t, err)
// 	assert.Equal(t, userUpdated.Name, userValues.Name)
// 	assert.Equal(t, userUpdated.Email, user.Email)

// 	t.Cleanup(func() {
// 		userDB.Delete(user.Email)
// 	})
// }

// func TestUpdateEmailUser(t *testing.T) {
// 	userDB := NewUser(test.Conn(t))
// 	user, _ := entity.NewUser("Paulo", "paulo@gmail.com", "123456")
// 	userDB.Create(user)
// 	userValues, _ := entity.NewUser("", "pedro@gmail.com", "")

// 	err := userDB.Update(userValues, user.ID.String())
// 	userUpdated, _ := userDB.FindByEmail(userValues.Email)

// 	assert.Nil(t, err)
// 	assert.Equal(t, userUpdated.Name, user.Name)
// 	assert.Equal(t, userUpdated.Email, userValues.Email)

// 	t.Cleanup(func() {
// 		userDB.Delete(userValues.Email)
// 	})
// }

// func TestUpdateUserNotFound(t *testing.T) {
// 	userDB := NewUser(test.Conn(t))
// 	userValues, _ := entity.NewUser("Pedro", "pedro@gmail.com", "")

// 	err := userDB.Update(userValues, "52fb6d91-29cb-43bb-965b-9dd37d482623")

// 	assert.Error(t, err, "not found")
// }

// func TestUpdateUserAlreadyRegistered(t *testing.T) {
// 	userDB := NewUser(test.Conn(t))
// 	user, _ := entity.NewUser("Paulo", "paulo@gmail.com", "123456")
// 	userDB.Create(user)

// 	userValues, _ := entity.NewUser("", "paulo@gmail.com", "")

// 	err := userDB.Update(userValues, user.ID.String())

// 	assert.Error(t, err, "invalid email")

// 	t.Cleanup(func() {
// 		userDB.Delete(user.Email)
// 	})
// }

// func TestListUserWithSuccess(t *testing.T) {
// 	userDB := NewUser(test.Conn(t))
// 	values := []struct {
// 		name  string
// 		email string
// 	}{
// 		{"Paulo", "paulo@gmail.com"},
// 		{"Pedro", "pedro@gmail.com"},
// 		{"maria", "maria@gmail.com"},
// 	}

// 	for _, v := range values {
// 		user, _ := entity.NewUser(v.name, v.email, "123465")

// 		userDB.Create(user)
// 	}

// 	users, err := userDB.List(0, 1)
// 	assert.Nil(t, err)
// 	assert.Len(t, users, 1)
// 	assert.Equal(t, users[0].Email, values[0].email)

// 	users, err = userDB.List(1, 2)
// 	assert.Nil(t, err)
// 	assert.Len(t, users, 2)
// 	assert.Equal(t, users[0].Email, values[1].email)
// 	assert.Equal(t, users[1].Email, values[2].email)

// 	t.Cleanup(func() {
// 		for _, v := range values {
// 			userDB.Delete(v.email)
// 		}
// 	})
// }

// func TestListUserWithInvalidOffset(t *testing.T) {
// 	userDB := NewUser(test.Conn(t))

// 	users, err := userDB.List(-1, 1)
// 	assert.Nil(t, users)
// 	assert.EqualError(t, err, "pq: OFFSET must not be negative")
// }
