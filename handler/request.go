package handler

import "github.com/labstack/echo/v4"

type createTweetRequest struct {
	Tweet string `json:"tweet" validate:"required,min=1,max=255"`
	Foo   string `json:"foo" validate:"required,min=1,max=255"`
}

func (r *createTweetRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}
