package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTweets_ツイート一覧が取得できること(t *testing.T) {
	setup()
	defer tearDown()

	testcases := []struct {
		name           string
		query          string
		goldenFilePath string
	}{
		{name: "クエリなし", query: "", goldenFilePath: "./testdata/tweet/tweets/no_query.golden"},
		{name: "limitで取得件数指定", query: "limit=2", goldenFilePath: "./testdata/tweet/tweets/limit.golden"},
		{name: "min_idで取得する最小のユーザIDを指定", query: "min_id=28", goldenFilePath: "./testdata/tweet/tweets/min_id.golden"},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			target := fmt.Sprintf("%v?%v", "/api/v1/tweets", tc.query)
			req := newRequest(http.MethodGet, target, nil)
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)

			assertResponse(t, rec.Result(), http.StatusOK, tc.goldenFilePath)
		})
	}
}

func TestNewTweets_ツイートを新規作成できること(t *testing.T) {
	setup()
	defer tearDown()

	type expected struct {
		httpStatusCode int
		goldenFilePath string
	}

	testcases := []struct {
		name     string
		body     string
		expected expected
	}{
		{
			name: "作成成功",
			body: `{"tweet": "foo"}`,
			expected: expected{
				httpStatusCode: http.StatusOK,
				goldenFilePath: "./testdata/tweet/new_tweet/success.golden",
			},
		},
		{
			name: "バリデーションエラー",
			body: "",
			expected: expected{
				httpStatusCode: http.StatusBadRequest,
				goldenFilePath: "./testdata/tweet/new_tweet/validation_error.golden",
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			req := newRequest(http.MethodPost, "/api/v1/tweets", strings.NewReader(tc.body))
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)

			assertResponse(t, rec.Result(), tc.expected.httpStatusCode, tc.expected.goldenFilePath)
		})
	}
}
