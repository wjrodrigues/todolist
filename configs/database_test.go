package configs

import (
	"testing"
	"todolist/pkg/test"

	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

func TestDBConnectionParamsIsValid(t *testing.T) {
	values, err := LoadEnv(".", test.FilePath(".env"))

	assert.Nil(t, err)
	assert.NotNil(t, DBConnection(values))
}

func TestDBConnectionParamsIsInValid(t *testing.T) {
	values, _ := LoadEnv(".", test.FilePath(".env"))

	values.DBDriver = ""

	assert.Panics(t, func() { DBConnection(values) })
}

func TestDBConnectionUserIsINvalid(t *testing.T) {
	values, _ := LoadEnv(".", test.FilePath(".env"))

	values.DBName = "any"

	assert.Panics(t, func() { DBConnection(values) })
}
