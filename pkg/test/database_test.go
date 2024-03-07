package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnSuccess(t *testing.T) {
	conn := Conn(t)

	assert.Nil(t, conn.Ping())
}
