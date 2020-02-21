package handler

import "github.com/labstack/echo/v4"

func (h *Handler) Register(v1 *echo.Group) {
	users := v1.Group("/users")
	users.GET("", h.Users)

	tweets := v1.Group("/tweets")
	tweets.GET("", h.Tweets)
	tweets.POST("", h.NewTweet)
	tweets.PUT("/:id/like", h.Like)
	tweets.PUT("/:id/unlike", h.Unlike)
}
