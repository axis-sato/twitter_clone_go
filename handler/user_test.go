package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUsers_ユーザ一覧が取得できること(t *testing.T) {
	setup()
	defer tearDown()

	testcases := []struct {
		name           string
		query          string
		goldenFilePath string
	}{
		{name: "クエリなし", query: "", goldenFilePath: "./testdata/user/users.golden"},
		{name: "limitで取得件数指定", query: "limit=2", goldenFilePath: "./testdata/user/users_with_limit.golden"},
		{name: "min_idで取得する最小のユーザIDを指定", query: "min_id=2", goldenFilePath: "./testdata/user/users_with_min_id.golden"},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			target := fmt.Sprintf("%v?%v", "/api/v1/users", tc.query)
			req := newRequest(http.MethodGet, target, nil)
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)

			assertResponse(t, rec.Result(), http.StatusOK, tc.goldenFilePath)
		})
	}
}
