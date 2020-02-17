package router

import (
	"github.com/c8112002/twitter_clone_go/handler"
	"github.com/c8112002/twitter_clone_go/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Recover())
	e.Validator = newValidator()
	e.HTTPErrorHandler = customHTTPErrorHandler
	return e
}

func customHTTPErrorHandler(err error, c echo.Context) {
	e := utils.NewError(err)
	res := handler.NewErrorResponse(e)
	c.JSON(e.HttpStatusCode, res)
}
