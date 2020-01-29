package handler

import (
	"math"
	"net/http"
	"strconv"

	"github.com/c8112002/twitter_clone_go/entities"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Tweets(c echo.Context) error {
	lastID, err := strconv.Atoi(c.QueryParam("last_id"))
	if err != nil {
		lastID = math.MaxInt64
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}

	tweets, err := h.tweetStore.FetchTweets(lastID, limit)
	if err != nil {
		c.Logger().Error("db error: " + err.Error())
	}

	firstTweet, err := h.tweetStore.FetchFirstTweet()

	if err != nil {
		c.Logger().Error("db error: " + err.Error())
	}

	res := new(tweetsResponse)
	for _, t := range *tweets {
		tr := newTweetResponse(t, t.IsLikedBy(entities.LoginUserID))
		res.Tweets = append(res.Tweets, tr)
	}

	res.ContainsFirstTweet = containsFirstTweet(firstTweet, tweets)

	return c.JSON(http.StatusOK, res)
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
