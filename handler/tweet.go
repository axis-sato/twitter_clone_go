package handler

import (
	"math"
	"net/http"
	"strconv"

	"github.com/c8112002/twitter_clone_go/entities"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Tweets(c echo.Context) error {
	//time.Sleep(1 * time.Second)
	tweets, err := h.tweetStore.FetchTweets(maxID(c), minID(c), limit(c))

	if err != nil {
		c.Logger().Error("db error: " + err.Error())
	}

	firstTweet, err := h.tweetStore.FetchFirstTweet()

	if err != nil {
		c.Logger().Error("db error: " + err.Error())
	}

	res := newEmptyTweetsResponse()
	for _, t := range *tweets {
		tr := newTweetResponse(t, t.IsLikedBy(entities.LoginUserID))
		res.Tweets = append(res.Tweets, tr)
	}

	res.ContainsFirstTweet = containsFirstTweet(firstTweet, tweets)

	return c.JSON(http.StatusOK, res)
}

func maxID(c echo.Context) int {
	mid, err := strconv.Atoi(c.QueryParam("max_id"))
	if err != nil {
		return math.MaxInt32
	}
	return mid
}

func minID(c echo.Context) int {
	mid, err := strconv.Atoi(c.QueryParam("min_id"))
	if err != nil {
		return 1
	}
	return mid
}

func limit(c echo.Context) int {
	l := c.QueryParam("limit")
	if l == "" {
		return math.MaxInt32
	}
	limit, err := strconv.Atoi(l)
	if err != nil {
		return math.MaxUint32
	}
	return limit
}

// tweetsにfirstTweetが含まれている場合trueを返す
func containsFirstTweet(firstTweet *entities.Tweet, tweets *entities.Tweets) bool {
	for _, t := range *tweets {
		if t.ID == firstTweet.ID {
			return true
		}
	}

	return false
}
