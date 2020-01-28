package handler

import (
	"math"
	"net/http"
	"strconv"

	"github.com/c8112002/twitter_clone_go/entities"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Users(c echo.Context) error {
	lastID, err := strconv.Atoi(c.QueryParam("last_id"))
	if err != nil {
		lastID = math.MaxInt64
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}

	users, err := h.userStore.FetchUsers(lastID, limit)
	if err != nil {
		c.Logger().Error("db error: " + err.Error())
	}

	res := new(usersResponse)
	for _, u := range users {
		ur := newUserResponse(u, u.IsFollowedBy(entities.LoginUserID))
		res.Users = append(res.Users, ur)
	}

	return c.JSON(http.StatusOK, res)
}
