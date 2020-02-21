package handler

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func tweetID(c echo.Context) (uint, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
