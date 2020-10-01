package login

import (
	"net/http"
	"strings"
	"testing"

	"github.com/evandroferreiras/gopher-city-bank/app/common/httputil"
	"github.com/evandroferreiras/gopher-city-bank/app/common/service"
	"github.com/evandroferreiras/gopher-city-bank/app/common/testutils"
	"github.com/evandroferreiras/gopher-city-bank/app/modules/login/mocks"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func setupService() *mocks.Service {
	return &mocks.Service{}
}

func Test_SignIn_ShouldReturnStatusOk_WhenGiveValidCpfAndSecret(t *testing.T) {
	serviceMock := setupService()
	serviceMock.On("SignIn", "12345612", "xxxxx").Return("someJwt", nil)
	reqJSON := `{
				  "cpf": "12345612",
				  "secret": "xxxxx"
				}
				`
	rec, ctx := testutils.GetRecordedAndContext(echo.POST, "/api/login", strings.NewReader(reqJSON))

	handler := Handler{LoginService: serviceMock}
	assert.NoError(t, handler.SignIn(ctx))

	assert.Equal(t, http.StatusOK, rec.Code)
	m := testutils.ResponseToMap(rec.Body.Bytes())
	assert.Equal(t, "someJwt", m["token"])
}

func Test_SignIn_ShouldReturnBadRequest_WhenHasErrorOnBody(t *testing.T) {
	serviceMock := setupService()

	reqJSON := `{
				 some invalid body
				`
	rec, ctx := testutils.GetRecordedAndContext(echo.POST, "/api/login", strings.NewReader(reqJSON))
	handler := Handler{LoginService: serviceMock}
	assert.NoError(t, handler.SignIn(ctx))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	m := testutils.ResponseToMap(rec.Body.Bytes())
	assert.Equal(t, httputil.HTTPErrorParseBody.Message, m["message"])
}

func Test_SignIn_ShouldReturnBadRequest_WhenBodyMissRequiredFields(t *testing.T) {
	serviceMock := setupService()

	reqJSON := `{}`

	rec, ctx := testutils.GetRecordedAndContext(echo.POST, "/api/login", strings.NewReader(reqJSON))
	handler := Handler{LoginService: serviceMock}
	assert.NoError(t, handler.SignIn(ctx))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	m := testutils.ResponseToMap(rec.Body.Bytes())
	assert.Equal(t, httputil.HTTPErrorValidateBody.Message, m["message"])
}

func Test_SignIn_ShouldReturnBadRequest_WhenGotGenericError(t *testing.T) {
	serviceMock := setupService()
	serviceMock.On("SignIn", "12345612", "xxxxx").Return("", errors.New("some error"))

	reqJSON := `{
				  "cpf": "12345612",
				  "secret": "xxxxx"
				}
				`
	rec, ctx := testutils.GetRecordedAndContext(echo.POST, "/api/login", strings.NewReader(reqJSON))
	handler := Handler{LoginService: serviceMock}
	assert.NoError(t, handler.SignIn(ctx))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	m := testutils.ResponseToMap(rec.Body.Bytes())
	assert.Equal(t, "some error", m["message"])
}

func Test_SignIn_ShouldReturnNotFound_WhenGotNotFoundError(t *testing.T) {
	serviceMock := setupService()
	serviceMock.On("SignIn", "12345612", "xxxxx").Return("", errors.Wrap(service.ErrorNotFound, "some error"))

	reqJSON := `{
				  "cpf": "12345612",
				  "secret": "xxxxx"
				}
				`
	rec, ctx := testutils.GetRecordedAndContext(echo.POST, "/api/login", strings.NewReader(reqJSON))
	handler := Handler{LoginService: serviceMock}
	assert.NoError(t, handler.SignIn(ctx))
	assert.Equal(t, http.StatusNotFound, rec.Code)
	m := testutils.ResponseToMap(rec.Body.Bytes())
	assert.Equal(t, "some error: not found", m["message"])
}
