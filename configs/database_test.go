package configs

import (
	"testing"
	"todolist/pkg/file"

	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

func TestDBConnectionParamsIsValid(t *testing.T) {
	values, err := LoadEnv(".", file.Path(".env"))

	assert.Nil(t, err)
	assert.NotNil(t, DBConnection(values))
}

func TestDBConnectionParamsIsInValid(t *testing.T) {
	values, _ := LoadEnv(".", file.Path(".env"))

	values.DBDriver = ""

	assert.Panics(t, func() { DBConnection(values) })
}

func TestDBConnectionUserIsINvalid(t *testing.T) {
	values, _ := LoadEnv(".", file.Path(".env"))

	values.DBName = "any"

	assert.Panics(t, func() { DBConnection(values) })
}
