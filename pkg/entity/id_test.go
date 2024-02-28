package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGenerateUUIDSuccess(t *testing.T) {
	id := NewID()

	_, err := uuid.Parse(id.String())

	assert.Nil(t, err)
}

func TestFailedParseUUID(t *testing.T) {
	_, err := ParseID("")

	assert.Error(t, err)
}
