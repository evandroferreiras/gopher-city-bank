package login

import (
	"net/http"

	"github.com/evandroferreiras/gopher-city-bank/app/common/httputil"
	"github.com/evandroferreiras/gopher-city-bank/app/representation"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Handler is a struct to Login handler
type Handler struct {
	LoginService Service
}

// NewHandler is a constructor to Login Handler
func NewHandler() *Handler {
	return &Handler{
		LoginService: NewService(),
	}
}

// SignIn godoc
// @Summary SignIn for existing user
// @Description SignIn for existing user
// @Tags login
// @Accept  json
// @Produce  json
// @Param user body representation.LoginBody true "Credentials to use"
// @Success 200 {object} representation.LoginResponse
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Router /api/login [post]
// SignIn for existing user
func (h Handler) SignIn(c echo.Context) error {
	login := &representation.LoginBody{}

	valid, err := httputil.IsValid(c, login)
	if err != nil || !valid {
		return err
	}

	jwtToken, err := h.LoginService.SignIn(login.Cpf, login.Secret)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusNotFound, httputil.NewError(http.StatusNotFound, err))
	}

	return c.JSON(http.StatusOK, representation.LoginResponse{Token: jwtToken})
}
