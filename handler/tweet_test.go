package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTweets_ユーザ一覧が取得できること(t *testing.T) {
	setup()
	defer tearDown()

	testcases := []struct {
		name           string
		query          string
		goldenFilePath string
	}{
		{name: "クエリなし", query: "", goldenFilePath: "./testdata/tweet/tweets.golden"},
		{name: "limitで取得件数指定", query: "limit=2", goldenFilePath: "./testdata/tweet/tweets_with_limit.golden"},
		{name: "min_idで取得する最小のユーザIDを指定", query: "min_id=28", goldenFilePath: "./testdata/tweet/tweets_with_min_id.golden"},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			target := fmt.Sprintf("%v?%v", "/api/v1/tweets", tc.query)
			req := newRequest(http.MethodGet, target, nil)
			rec := httptest.NewRecorder()
			c := newEchoContext(req, rec)

			assert.NoError(t, h.Tweets(c))
			assertResponse(t, rec.Result(), 200, tc.goldenFilePath)
		})
	}
}
