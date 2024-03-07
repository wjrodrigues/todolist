package configs

import (
	"testing"
	"todolist/pkg/file"

	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {
	values, err := LoadEnv(".", file.Path(".env"))

	assert.Nil(t, err)
	assert.NotEmpty(t, values)
}

func TestFileNotExists(t *testing.T) {
	values, err := LoadEnv(".", "any")

	assert.NotNil(t, err)
	assert.Equal(t, values, env{})
}
