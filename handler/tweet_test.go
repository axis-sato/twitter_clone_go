package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/c8112002/twitter_clone_go/models"
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

	type expected struct {
		httpStatusCode int
		goldenFilePath string
		tweetCount     int
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
				tweetCount:     41,
			},
		},
		{
			name: "バリデーションエラー",
			body: "",
			expected: expected{
				httpStatusCode: http.StatusBadRequest,
				goldenFilePath: "./testdata/tweet/new_tweet/validation_error.golden",
				tweetCount:     40,
			},
		},
	}

	for _, tc := range testcases {
		setup()

		t.Run(tc.name, func(t *testing.T) {
			req := newRequest(http.MethodPost, "/api/v1/tweets", strings.NewReader(tc.body))
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)

			assertResponse(t, rec.Result(), tc.expected.httpStatusCode, tc.expected.goldenFilePath)

			count, _ := models.Tweets().Count(ctx, d)
			assert.Equal(t, tc.expected.tweetCount, int(count))
		})

		tearDown()
	}
}
