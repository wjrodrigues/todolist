package item

import (
	"testing"
	"time"
	"todolist/pkg/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewItem(t *testing.T) {
	item := NewItem("Pay bankslip", "any", PENDING, entity.NewID())

	assert.NotNil(t, item)
	assert.NotEmpty(t, item.ID)
	assert.Equal(t, item.Title, "Pay bankslip")
	assert.Equal(t, item.Description, "any")
	assert.Equal(t, time.Now().Format(time.RFC822), item.CreatedAt.Format(time.RFC822))
	assert.Equal(t, time.Now().Format(time.RFC822), item.UpdatedAt.Format(time.RFC822))
}

func TestChangeStatus(t *testing.T) {
	item := NewItem("Pay bankslip", "any", PENDING, entity.NewID())

	item.Canceled()
	assert.Equal(t, item.Status, CANCELED)

	item.Pending()
	assert.Equal(t, item.Status, PENDING)

	item.InProgress()
	assert.Equal(t, item.Status, IN_PROGRESS)

	item.Completed()
	assert.Equal(t, item.Status, COMPLETED)
}
