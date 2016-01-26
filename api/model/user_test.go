package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user := NewUser(1, "hoge")

	assert.NotEmpty(t, user)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "hoge", user.Name)
}
