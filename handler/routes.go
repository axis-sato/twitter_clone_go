package handler

import "github.com/labstack/echo/v4"

func (h *Handler) Register(v1 *echo.Group) {
	users := v1.Group("/users")
	users.GET("", h.Users)
}