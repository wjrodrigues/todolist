package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	item := NewItem("Pay bankslip", "any item", "pending")
	items := []Item{*item}
	owner, _ := NewUser("Pedro", "pedro@email.com", "123")
	list := NewList("Pay bankslips", "any list", "in_progress", items, *owner)

	assert.NotNil(t, list)
	assert.NotEmpty(t, list.ID)
	assert.Equal(t, list.Title, "Pay bankslips")
	assert.Equal(t, list.Description, "any list")
	assert.Equal(t, list.Status, "in_progress")
	assert.Equal(t, list.Items, items)
	assert.Equal(t, time.Now().Format(time.RFC822), item.CreatedAt.Format(time.RFC822))
	assert.Equal(t, time.Now().Format(time.RFC822), item.UpdatedAt.Format(time.RFC822))
}
