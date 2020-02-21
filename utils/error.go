package utils

import (
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
)

type InvalidParamError struct {
	Message string
}

func (e *InvalidParamError) Error() string {
	return e.Message
}

type Error struct {
	Error          string
	HttpStatusCode int
}

func NewError(err error) Error {
	switch e := err.(type) {
	case *echo.HTTPError:
		return Error{Error: e.Message.(string), HttpStatusCode: e.Code}
	case validator.ValidationErrors:
		return Error{
			Error:          e.Error(),
			HttpStatusCode: http.StatusBadRequest,
		}
	case *InvalidParamError:
		return Error{
			Error:          e.Error(),
			HttpStatusCode: http.StatusBadRequest,
		}
	default:
		return newInternalServerError()
	}
}

func newInternalServerError() Error {
	return Error{Error: "Internal Server Error", HttpStatusCode: http.StatusInternalServerError}
}
