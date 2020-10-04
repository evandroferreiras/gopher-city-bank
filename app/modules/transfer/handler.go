package transfer

import (
	"net/http"
	"strings"

	"github.com/evandroferreiras/gopher-city-bank/app/common/httputil"
	"github.com/evandroferreiras/gopher-city-bank/app/common/jwt"
	"github.com/evandroferreiras/gopher-city-bank/app/common/service"
	"github.com/evandroferreiras/gopher-city-bank/app/representation"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const emptyAccountID = ""

// Handler is a struct to Transfer handler
type Handler struct {
	TransferService Service
}

// NewHandler is a constructor to Transfer Handler
func NewHandler() *Handler {
	return &Handler{
		TransferService: NewService(),
	}
}

// TransferToAnotherAccount godoc
// @Summary Transfer money to another account
// @Description Transfer money to another account
// @Tags transfer
// @Accept  json
// @Produce json
// @Param user body representation.TransferBody true "account destination and amount"
// @Success 200 {object} representation.AccountBalanceResponse
// @Failure 400 {object} httputil.HTTPError
// @Failure 401 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Security ApiKeyAuth
// @Router /api/transfers [post]
func (h Handler) TransferToAnotherAccount(c echo.Context) error {
	accountOriginID, err := getAccountIDFromHeader(c)
	if err != nil || accountOriginID == emptyAccountID {
		return err
	}

	transferBody := &representation.TransferBody{}

	valid, err := httputil.IsValid(c, transferBody)
	if err != nil || !valid {
		return err
	}

	account, err := h.TransferService.TransferBetweenAccount(accountOriginID, transferBody.AccountDestinationID, transferBody.Amount)
	if err != nil {
		logrus.Error(err)

		if errors.Cause(err) == service.ErrorNotFound {
			return c.JSON(http.StatusNotFound, httputil.NewError(http.StatusNotFound, err))
		}
		return c.JSON(http.StatusBadRequest, httputil.NewError(http.StatusBadRequest, err))
	}

	return c.JSON(http.StatusOK, representation.ModelToAccountBalanceResponse(account))
}

// List godoc
// @Summary List all transfers of an account
// @Description List all transfers of an account
// @Tags transfer
// @Produce json
// @Success 200 {object} representation.TransferListResponse
// @Failure 400 {object} httputil.HTTPError
// @Failure 401 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Security ApiKeyAuth
// @Router /api/transfers [get]
func (h Handler) List(c echo.Context) error {
	accountOriginID, err := getAccountIDFromHeader(c)
	if err != nil || accountOriginID == emptyAccountID {
		return err
	}

	withdraws, err := h.TransferService.GetAllWithdrawsOf(accountOriginID)
	if err != nil {
		logrus.Error(err)
		if errors.Cause(err) == service.ErrorNotFound {
			return c.JSON(http.StatusNotFound, httputil.NewError(http.StatusNotFound, err))
		}
		return c.JSON(http.StatusBadRequest, httputil.NewError(http.StatusBadRequest, err))
	}

	deposits, err := h.TransferService.GetAllDepositsTo(accountOriginID)
	if err != nil {
		logrus.Error(err)
		if errors.Cause(err) == service.ErrorNotFound {
			return c.JSON(http.StatusNotFound, httputil.NewError(http.StatusNotFound, err))
		}
		return c.JSON(http.StatusBadRequest, httputil.NewError(http.StatusBadRequest, err))
	}

	return c.JSON(http.StatusOK, representation.NewTransferListResponse(withdraws, deposits))
}

func getAccountIDFromHeader(c echo.Context) (string, error) {
	jwtToken := ""
	// Format expected: Token JwtToken
	authHeaderSplit := strings.Split(c.Request().Header.Get("Authorization"), " ")
	if len(authHeaderSplit) == 2 {
		jwtToken = authHeaderSplit[1]
	}

	if jwtToken == "" {
		return emptyAccountID, c.JSON(http.StatusUnauthorized, httputil.HTTPErrorInvalidJWT)
	}

	accountID, err := jwt.GetIDFromJWT(jwtToken)
	if err != nil {
		logrus.Error(errors.Wrap(err, "error when trying to get ID from JWT"))
		return emptyAccountID, c.JSON(http.StatusUnauthorized, httputil.HTTPErrorInvalidJWT)
	}

	return accountID, nil
}
