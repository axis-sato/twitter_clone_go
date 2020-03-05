package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Recover())
	e.Validator = newValidator()
	return e
}
