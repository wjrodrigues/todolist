package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMehtodsResponse(t *testing.T) {
	response := Response[string]{}

	assert.False(t, response.HasError())

	response.AddError(errors.New("test"))

	assert.True(t, response.HasError())

	response.AddResult("any")

	assert.Equal(t, response.Result(), "any")
}
