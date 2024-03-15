package list

import (
	"testing"
	"time"
	"todolist/internal/entity/item"
	"todolist/internal/entity/user"
	"todolist/pkg/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	owner, _ := user.NewUser("Pedro", "pedro@email.com", "123")
	list := NewList("Pay bankslips", "any list", IN_PROGRESS, *owner)

	assert.NotNil(t, list)
	assert.NotEmpty(t, list.ID)
	assert.Equal(t, list.Title, "Pay bankslips")
	assert.Equal(t, list.Description, "any list")
	assert.Equal(t, list.Status, IN_PROGRESS)
}

func TestAddItem(t *testing.T) {
	owner, _ := user.NewUser("Pedro", "pedro@email.com", "123")
	list := NewList("Pay bankslips", "any list", IN_PROGRESS, *owner)
	itemInstance := item.NewItem("Pay bankslip", "any item", item.PENDING, entity.NewID())

	list.AddItem(*itemInstance)

	assert.Equal(t, time.Now().Format(time.RFC822), itemInstance.CreatedAt.Format(time.RFC822))
	assert.Equal(t, time.Now().Format(time.RFC822), itemInstance.UpdatedAt.Format(time.RFC822))
	assert.Equal(t, list.Items, []*item.Item{itemInstance})
}

func TestChangeStatus(t *testing.T) {
	owner, _ := user.NewUser("Pedro", "pedro@email.com", "123")
	list := NewList("Pay bankslips", "any list", PENDING, *owner)
	itemInstance := item.NewItem("Pay bankslip", "any item", item.PENDING, entity.NewID())
	list.AddItem(*itemInstance)

	itemInstance = item.NewItem("Pay bankslip", "any item", item.PENDING, entity.NewID())
	list.AddItem(*itemInstance)

	list.Canceled()
	assert.Equal(t, list.Status, CANCELED)
	for _, item := range list.Items {
		assert.Equal(t, item.Status, CANCELED)
	}

	list.Pending()
	assert.Equal(t, list.Status, PENDING)
	for _, item := range list.Items {
		assert.Equal(t, item.Status, PENDING)
	}

	list.InProgress()
	assert.Equal(t, list.Status, IN_PROGRESS)
	for _, item := range list.Items {
		assert.Equal(t, item.Status, IN_PROGRESS)
	}

	list.Completed()
	assert.Equal(t, list.Status, COMPLETED)
	for _, item := range list.Items {
		assert.Equal(t, item.Status, COMPLETED)
	}
}
