package account

import (
	"net/http"
	"strings"
	"testing"

	"github.com/evandroferreiras/gopher-city-bank/app/common/service"
	"github.com/pkg/errors"

	"github.com/evandroferreiras/gopher-city-bank/app/model"

	"github.com/evandroferreiras/gopher-city-bank/app/common/httputil"
	"github.com/evandroferreiras/gopher-city-bank/app/common/testutils"
	"github.com/evandroferreiras/gopher-city-bank/app/modules/account/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupService() *mocks.Service {
	return &mocks.Service{}
}

func Test_CreateAccount_ShouldReturnStatusCreated_WhenCreateWithSuccess(t *testing.T) {
	serviceMock := setupService()
	account := &model.Account{
		ID:      "1",
		Name:    "Bruce Wayne",
		Cpf:     "12345612",
		Balance: 1000000,
	}
	serviceMock.On("Create", mock.Anything).Return(account, nil)

	reqJSON := `{
				  "balance": 1000000,
				  "cpf": "12345612",
				  "name": "Bruce Wayne",
				  "secret": "xxxxx"
				}
				`
	rec, ctx := testutils.GetRecordedAndContext(echo.POST, "/api/accounts", strings.NewReader(reqJSON))

	handler := Handler{AccountService: serviceMock}
	assert.NoError(t, handler.CreateAccount(ctx))
	t.Log(rec.Body.String())
	if assert.Equal(t, http.StatusCreated, rec.Code) {
		m := testutils.ResponseToMap(rec.Body.Bytes())
		assert.Equal(t, "1", m["id"])
		assert.Equal(t, "Bruce Wayne", m["name"])
	}
}

func Test_CreateAccount_ShouldReturnBadRequest_WhenHasErrorOnBody(t *testing.T) {
	serviceMock := setupService()
	serviceMock.On("Create", mock.Anything).Return(nil, nil)

	reqJSON := `{
				 some invalid body
				`
	rec, ctx := testutils.GetRecordedAndContext(echo.POST, "/api/accounts", strings.NewReader(reqJSON))
	handler := Handler{AccountService: serviceMock}
	assert.NoError(t, handler.CreateAccount(ctx))

	t.Log(rec.Body.String())

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	m := testutils.ResponseToMap(rec.Body.Bytes())
	assert.Equal(t, httputil.HTTPErrorParseBody.Message, m["message"])
}

func Test_CreateAccount_ShouldReturnBadRequest_WhenBodyMissRequiredFields(t *testing.T) {
	serviceMock := setupService()
	serviceMock.On("Create", mock.Anything).Return(nil, nil)

	reqJSON := `{}`
	rec, ctx := testutils.GetRecordedAndContext(echo.POST, "/api/accounts", strings.NewReader(reqJSON))
	handler := Handler{AccountService: serviceMock}
	assert.NoError(t, handler.CreateAccount(ctx))

	t.Log(rec.Body.String())

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	m := testutils.ResponseToMap(rec.Body.Bytes())
	assert.Equal(t, httputil.HTTPErrorValidateBody.Message, m["message"])
}

func Test_CreateAccount_ShouldReturnBadRequest_WhenCantCreateAnAccount(t *testing.T) {
	serviceMock := setupService()
	serviceMock.On("Create", mock.Anything).Return(nil, errors.New("some error"))

	reqJSON := `{
				  "balance": 1000000,
				  "cpf": "12345612",
				  "name": "Bruce Wayne",
				  "secret": "xxxxx"
				}
				`
	rec, ctx := testutils.GetRecordedAndContext(echo.POST, "/api/accounts", strings.NewReader(reqJSON))
	handler := Handler{AccountService: serviceMock}
	assert.NoError(t, handler.CreateAccount(ctx))

	t.Log(rec.Body.String())

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	m := testutils.ResponseToMap(rec.Body.Bytes())
	assert.Equal(t, "some error", m["message"])
}

func Test_CreateAccount_ShouldReturnBadRequest_WhenBalanceIsLessThanZero(t *testing.T) {
	serviceMock := setupService()
	serviceMock.On("Create", mock.Anything).Return(nil, errors.New("some error"))

	reqJSON := `{
				  "balance": -1,
				  "cpf": "12345612",
				  "name": "Bruce Wayne",
				  "secret": "xxxxx"
				}
				`
	rec, ctx := testutils.GetRecordedAndContext(echo.POST, "/api/accounts", strings.NewReader(reqJSON))
	handler := Handler{AccountService: serviceMock}
	assert.NoError(t, handler.CreateAccount(ctx))

	t.Log(rec.Body.String())

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	m := testutils.ResponseToMap(rec.Body.Bytes())

	assert.Equal(t, httputil.HTTPErrorValidateBody.Message, m["message"])

}

func Test_GetAllAccounts_ShouldReturnStatusOk_WhenReturnWithSuccess(t *testing.T) {
	serviceMock := setupService()
	serviceMock.On("GetAccounts").Return([]model.Account{{ID: "1"}, {ID: "2"}}, nil)
	rec, ctx := testutils.GetRecordedAndContext(echo.GET, "/api/accounts", nil)
	handler := Handler{AccountService: serviceMock}
	assert.NoError(t, handler.GetAllAccounts(ctx))

	t.Log(rec.Body.String())
	assert.Equal(t, http.StatusOK, rec.Code)
	m := testutils.ResponseToMap(rec.Body.Bytes())
	assert.Equal(t, "1", m["accounts"].([]interface{})[0].(map[string]interface{})["id"])
	assert.Equal(t, "2", m["accounts"].([]interface{})[1].(map[string]interface{})["id"])
	assert.Equal(t, float64(2), m["count"])
}

func Test_GetAllAccounts_ShouldReturnBadRequest_WhenGotError(t *testing.T) {
	serviceMock := setupService()
	serviceMock.On("GetAccounts").Return(nil, errors.New("some error"))
	rec, ctx := testutils.GetRecordedAndContext(echo.GET, "/api/accounts", nil)
	handler := Handler{AccountService: serviceMock}

	assert.NoError(t, handler.GetAllAccounts(ctx))
	m := testutils.ResponseToMap(rec.Body.Bytes())
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "some error", m["message"])
}

func Test_GetAccountBalance_ShouldReturnStatusOk_WhenReturnWithSuccess(t *testing.T) {
	accountID := "1"

	serviceMock := setupService()
	serviceMock.On("GetAccount", accountID).Return(&model.Account{ID: accountID, Name: "Bruce Wayne", Balance: 10000}, nil)
	rec, ctx := testutils.GetRecordedAndContext(echo.GET, "/api/accounts/:account_id/balance", nil)
	ctx.SetParamNames("account_id")
	ctx.SetParamValues(accountID)

	handler := Handler{AccountService: serviceMock}
	assert.NoError(t, handler.GetAccountBalance(ctx))
	t.Log(rec.Body.String())
	m := testutils.ResponseToMap(rec.Body.Bytes())
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, float64(10000), m["balance"])

}

func Test_GetAccountBalance_ShouldReturnBadRequest_WhenGotError(t *testing.T) {
	accountID := "1"

	serviceMock := setupService()
	serviceMock.On("GetAccount", accountID).Return(nil, errors.New("some error"))
	rec, ctx := testutils.GetRecordedAndContext(echo.GET, "/api/accounts/:account_id/balance", nil)
	ctx.SetParamNames("account_id")
	ctx.SetParamValues(accountID)

	handler := Handler{AccountService: serviceMock}
	assert.NoError(t, handler.GetAccountBalance(ctx))
	t.Log(rec.Body.String())
	m := testutils.ResponseToMap(rec.Body.Bytes())
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	assert.Equal(t, "some error", m["message"])
}

func Test_GetAccountBalance_ShouldReturnNotFound_WhenGotNotFoundError(t *testing.T) {
	accountID := "1"

	serviceMock := setupService()
	serviceMock.On("GetAccount", accountID).Return(nil, errors.Wrap(service.ErrorNotFound, "some error"))
	rec, ctx := testutils.GetRecordedAndContext(echo.GET, "/api/accounts/:account_id/balance", nil)
	ctx.SetParamNames("account_id")
	ctx.SetParamValues(accountID)

	handler := Handler{AccountService: serviceMock}
	assert.NoError(t, handler.GetAccountBalance(ctx))
	t.Log(rec.Body.String())

	assert.Equal(t, http.StatusNotFound, rec.Code)
	m := testutils.ResponseToMap(rec.Body.Bytes())
	assert.Equal(t, "some error: not found", m["message"])
}
