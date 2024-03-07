package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMountFullPath(t *testing.T) {
	path := Path(".env")

	assert.Equal(t, path, "/app/.env")
}
