package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewItem(t *testing.T) {
	item := NewItem("Pay bankslip", "any", "pending")

	assert.NotNil(t, item)
	assert.NotEmpty(t, item.ID)
	assert.Equal(t, item.Title, "Pay bankslip")
	assert.Equal(t, item.Description, "any")
	assert.Equal(t, time.Now().Format(time.RFC822), item.CreatedAt.Format(time.RFC822))
	assert.Equal(t, time.Now().Format(time.RFC822), item.UpdatedAt.Format(time.RFC822))
}
