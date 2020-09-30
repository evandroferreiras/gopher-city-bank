package account

import (
	"net/http"

	"github.com/evandroferreiras/gopher-city-bank/app/common/httputil"
	"github.com/evandroferreiras/gopher-city-bank/app/model"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Handler is a struct to Account handler
type Handler struct {
	AccountService Service
}

// NewHandler is a constructor to Account Handler
func NewHandler() *Handler {
	return &Handler{
		AccountService: NewService(),
	}
}

// CreateAccount godoc
// @Summary Create account
// @Description Creates a new account
// @tags accounts
// @Accept  json
// @Produce  json
// @Param account body model.NewAccount true "Add account"
// @Success 201
// @Failure 400 {object} httputil.HTTPError
// @Router /api/accounts [post]
// CreateAccount is a creator an account
func (h *Handler) CreateAccount(c echo.Context) error {
	account := &model.NewAccount{}

	if err := c.Bind(account); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, httputil.HTTPErrorParseBody)
	}

	if err := c.Validate(account); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, httputil.NewHTTPErrorValidateBody(err))
	}

	newAccount, err := h.AccountService.Create(account)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, httputil.NewError(http.StatusBadRequest, err))
	}
	return c.JSON(http.StatusCreated, newAccount)
}
