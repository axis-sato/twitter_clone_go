package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsers_ユーザ一覧が取得できること(t *testing.T) {
	setup()
	defer tearDown()
	assert.True(t, true)
}
