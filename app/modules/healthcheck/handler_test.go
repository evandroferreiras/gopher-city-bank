package healthcheck

import (
	"net/http"
	"testing"

	"github.com/evandroferreiras/gopher-city-bank/app/common/testutils"

	"github.com/evandroferreiras/gopher-city-bank/app/modules/healthcheck/mocks"
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

	rec, ctx := testutils.GetRecordedAndContext(echo.GET, "/api/healthcheck", nil)

	assert.NoError(t, handler.IsOnline(ctx))
	if assert.Equal(t, http.StatusOK, rec.Code) {
		assert.Equal(t, "WORKING", rec.Body.String())
	}
}

func Test_IsOnline_ShouldReturnNotWorking_IfServiceIsNotWorking(t *testing.T) {
	serviceMock := setupService()

	serviceMock.On("IsWorking").Return(false)
	handler := Handler{healthCheckService: serviceMock}

	rec, ctx := testutils.GetRecordedAndContext(echo.GET, "/api/healthcheck", nil)

	assert.NoError(t, handler.IsOnline(ctx))
	if assert.Equal(t, http.StatusBadRequest, rec.Code) {
		assert.Equal(t, "NOT WORKING", rec.Body.String())
	}
}
