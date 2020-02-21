package handler

import (
	"net/http"

	"github.com/c8112002/twitter_clone_go/utils"

	"github.com/c8112002/twitter_clone_go/entities"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Tweets(c echo.Context) error {
	//time.Sleep(1 * time.Second)
	tweets, err := h.tweetStore.FetchTweets(maxID(c), minID(c), limit(c))

	if err != nil {
		c.Logger().Error("db error: " + err.Error())
		return err
	}

	firstTweet, err := h.tweetStore.FetchFirstTweet()

	if err != nil {
		c.Logger().Error("db error: " + err.Error())
		return err
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
		return err
	}

	t, err := h.tweetStore.CreateTweet(r.Tweet, entities.LoginUserID)
	if err != nil {
		c.Logger().Error("db error: " + err.Error())
		return err
	}

	tr := newTweetResponse(t, t.IsLikedBy(entities.LoginUserID))

	return c.JSON(http.StatusOK, tr)
}

func (h *Handler) Like(c echo.Context) error {
	id, err := tweetID(c)
	if err != nil {
		c.Logger().Error("param error: " + err.Error())
		return &utils.InvalidParamError{Message: "Invalid Tweet ID"}
	}

	t, err := h.tweetStore.FindTweet(id)
	if err != nil {
		c.Logger().Error("db error: " + err.Error())
		return err
	}

	if t.IsLikedBy(entities.LoginUserID) {
		tr := newTweetResponse(t, t.IsLikedBy(entities.LoginUserID))
		return c.JSON(http.StatusOK, tr)
	}

	tweet, err := h.tweetStore.Like(t, entities.LoginUserID)
	if err != nil {
		c.Logger().Error("db error: " + err.Error())
		return err
	}
	tr := newTweetResponse(tweet, tweet.IsLikedBy(entities.LoginUserID))
	return c.JSON(http.StatusOK, tr)
}

func (h *Handler) Unlike(c echo.Context) error {
	id, err := tweetID(c)
	if err != nil {
		c.Logger().Error("param error: " + err.Error())
		return &utils.InvalidParamError{Message: "Invalid Tweet ID"}
	}

	t, err := h.tweetStore.FindTweet(id)
	if err != nil {
		c.Logger().Error("db error: " + err.Error())
		return err
	}

	if !t.IsLikedBy(entities.LoginUserID) {
		tr := newTweetResponse(t, t.IsLikedBy(entities.LoginUserID))
		return c.JSON(http.StatusOK, tr)
	}

	tweet, err := h.tweetStore.Unlike(t, entities.LoginUserID)
	if err != nil {
		c.Logger().Error("db error: " + err.Error())
		return err
	}
	tr := newTweetResponse(tweet, tweet.IsLikedBy(entities.LoginUserID))
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
