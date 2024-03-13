package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	owner, _ := NewUser("Pedro", "pedro@email.com", "123")
	list := NewList("Pay bankslips", "any list", "in_progress", *owner)

	assert.NotNil(t, list)
	assert.NotEmpty(t, list.ID)
	assert.Equal(t, list.Title, "Pay bankslips")
	assert.Equal(t, list.Description, "any list")
	assert.Equal(t, list.Status, "in_progress")
}

func TestAddItem(t *testing.T) {
	owner, _ := NewUser("Pedro", "pedro@email.com", "123")
	list := NewList("Pay bankslips", "any list", "in_progress", *owner)
	item := NewItem("Pay bankslip", "any item", "pending", List{})

	list.AddItem(*item)

	assert.Equal(t, time.Now().Format(time.RFC822), item.CreatedAt.Format(time.RFC822))
	assert.Equal(t, time.Now().Format(time.RFC822), item.UpdatedAt.Format(time.RFC822))
	assert.Equal(t, list.Items, []Item{*item})
}
