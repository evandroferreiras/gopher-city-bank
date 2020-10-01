package httputil

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// IsValid decodes and validates a body based on a struct
func IsValid(c echo.Context, i interface{}) (valid bool, err error) {
	if err := c.Bind(i); err != nil {
		logrus.Error(err)
		return false, c.JSON(http.StatusBadRequest, HTTPErrorParseBody)
	}

	if err := c.Validate(i); err != nil {
		logrus.Error(err)
		return false, c.JSON(http.StatusBadRequest, NewHTTPErrorValidateBody(err))
	}
	return true, nil
}
