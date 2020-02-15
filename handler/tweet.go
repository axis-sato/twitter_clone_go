package handler

import (
	"net/http"

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

func (h *Handler) NewTweet(c echo.Context) error {
	r := new(createTweetRequest)
	if err := r.bind(c); err != nil {
		c.Logger().Error("request error: " + err.Error())
		// TODO適切なエラーレスポンスを返す
		return c.JSON(http.StatusBadRequest, "validation error")
	}

	t, err := h.tweetStore.CreateTweet(r.Tweet, entities.LoginUserID)
	if err != nil {
		c.Logger().Error("db error: " + err.Error())
	}

	tr := newTweetResponse(t, t.IsLikedBy(entities.LoginUserID))

	return c.JSON(http.StatusOK, tr)
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
