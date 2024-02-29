package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMountFullPath(t *testing.T) {
	path := FilePath(".env")

	assert.Equal(t, path, "//app/.env")
}
