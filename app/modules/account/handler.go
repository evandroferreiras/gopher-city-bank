package account

import (
	"net/http"

	"github.com/evandroferreiras/gopher-city-bank/app/model"

	"github.com/evandroferreiras/gopher-city-bank/app/representation"

	"github.com/evandroferreiras/gopher-city-bank/app/common/httputil"
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
// @Param account body representation.NewAccountBody true "Add account"
// @Success 201 {object} representation.AccountResponse
// @Failure 400 {object} httputil.HTTPError
// @Router /api/accounts [post]
// CreateAccount is a creator an account
func (h *Handler) CreateAccount(c echo.Context) error {
	account := &representation.NewAccountBody{}

	if err := c.Bind(account); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, httputil.HTTPErrorParseBody)
	}

	if err := c.Validate(account); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, httputil.NewHTTPErrorValidateBody(err))
	}

	createdAccount, err := h.AccountService.Create(account.ToModel())
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, httputil.NewError(http.StatusBadRequest, err))
	}
	return c.JSON(http.StatusCreated, representation.ModelToAccountResponse(createdAccount))
}

// GetAllAccounts godoc
// @Summary Get all accounts
// @Tags accounts
// @Accept  json
// @Produce  json
// @Success 200 {object} representation.AccountsList
// @Failure 400 {object} httputil.HTTPError
// @Router /api/accounts [get]
// GetAllAccounts returns all accounts
func (h *Handler) GetAllAccounts(c echo.Context) error {
	accounts, err := h.AccountService.GetAccounts()
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, httputil.NewError(http.StatusBadRequest, err))
	}
	return c.JSON(http.StatusOK, getAccountsResponse(accounts))
}

func getAccountsResponse(accounts []model.Account) representation.AccountsList {
	accountListResponse := representation.AccountsList{}
	accountListResponse.Accounts = make([]representation.AccountResponse, 0)
	for _, account := range accounts {
		accountListResponse.Accounts = append(accountListResponse.Accounts, *representation.ModelToAccountResponse(&account))
	}
	accountListResponse.Count = len(accountListResponse.Accounts)

	return accountListResponse
}
