package account

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/evandroferreiras/gopher-city-bank/app/common/httputil"
	"github.com/evandroferreiras/gopher-city-bank/app/common/testutils"
	"github.com/evandroferreiras/gopher-city-bank/app/model"
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
	account := &model.AccountCreated{
		ID:      1,
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
		assert.Equal(t, float64(1), m["id"])
		assert.Equal(t, "Bruce Wayne", m["name"])
		assert.Equal(t, float64(1000000), m["balance"])
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
