package handler

import (
	"math"
	"strconv"

	"github.com/labstack/echo/v4"
)

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
