package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsers_ユーザ一覧が取得できること(t *testing.T) {
	setup()
	defer tearDown()

	req := newRequest(http.MethodGet, "/api/v1/users", nil)
	rec := httptest.NewRecorder()
	c := newEchoContext(req, rec)

	assert.NoError(t, h.Users(c))
	assertResponse(t, rec.Result(), 200, "./testdata/user/users.golden")
	var r usersResponse
	err := json.Unmarshal(rec.Body.Bytes(), &r)
	assert.NoError(t, err)
}
