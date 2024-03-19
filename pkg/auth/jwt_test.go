package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateJwt(t *testing.T) {
	jwtBody := JwtBody{ID: "123"}

	result, err := GenerateJWT(jwtBody, "123")
	assert.NotEmpty(t, result)
	assert.Nil(t, err)
}
