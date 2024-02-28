package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	Main()

	assert.Equal(t, true, true)
}
