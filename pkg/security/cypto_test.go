package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateFromPasswordSuccess(t *testing.T) {
	password := "123456"

	hash, err := GenerateFromPassword(password)

	assert.Nil(t, err)
	assert.NotEmpty(t, hash)
}

func TestGenerateFromPasswordFailedWhenLongPassword(t *testing.T) {
	var password = "LoremIpsumissimplydummytextoftheprintingandtypesettingindustryLoremIpsuab"

	hash, err := GenerateFromPassword(password)

	assert.Error(t, err)
	assert.Empty(t, hash)
}
