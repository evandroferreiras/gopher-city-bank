package healthcheck

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/evandroferreiras/gopher-city-bank/app/healthcheck/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupService() *mocks.Service {
	return &mocks.Service{}
}

func Test_IsOnline_ShouldReturnWorking_IfServiceIsWorking(t *testing.T) {
	serviceMock := setupService()

	serviceMock.On("IsWorking").Return(true)
	handler := Handler{healthCheckService: serviceMock}

	rec, ctx := getRecordedAndContext(echo.GET, "/api/healthcheck", nil)

	assert.NoError(t, handler.IsOnline(ctx))
	if assert.Equal(t, http.StatusOK, rec.Code) {
		assert.Equal(t, "WORKING", rec.Body.String())
	}
}

func Test_IsOnline_ShouldReturnNotWorking_IfServiceIsNotWorking(t *testing.T) {
	serviceMock := setupService()

	serviceMock.On("IsWorking").Return(false)
	handler := Handler{healthCheckService: serviceMock}

	rec, ctx := getRecordedAndContext(echo.GET, "/api/healthcheck", nil)

	assert.NoError(t, handler.IsOnline(ctx))
	if assert.Equal(t, http.StatusBadRequest, rec.Code) {
		assert.Equal(t, "NOT WORKING", rec.Body.String())
	}
}

func getRecordedAndContext(method string, target string, body io.Reader) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	req := httptest.NewRequest(method, target, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return rec, c
}
