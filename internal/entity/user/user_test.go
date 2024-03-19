package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "john@gmail.com", "123456")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@gmail.com", user.Email)

	_, err = NewUser(
		"John Doe",
		"john@gmail.com",
		"49d7CSQKfiCYbBR036kHWLmUZojK2m2L3348swmZCaH9de092D4OUrgoUD6jmeV21VVjldqPasd")
	assert.NotNil(t, err)
	assert.Error(t, err, "bcrypt: password length exceeds 72 bytes")
}

func TestValidatePassword(t *testing.T) {
	user, _ := NewUser("John Doe", "john@gmail.com", "123456")

	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("123"))
}
