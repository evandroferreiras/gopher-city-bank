package transfer

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/evandroferreiras/gopher-city-bank/app/common/httputil"
	"github.com/evandroferreiras/gopher-city-bank/app/common/service"
	"github.com/evandroferreiras/gopher-city-bank/app/common/testutils"
	"github.com/evandroferreiras/gopher-city-bank/app/modules/transfer/mocks"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func setupService() *mocks.Service {
	return &mocks.Service{}
}

func Test_TransferToAnotherAccount_ShouldReturnStatusOK_WhenUserIsAuthenticatedAndCanTransferWithSuccess(t *testing.T) {
	serviceMock := setupService()
	serviceMock.On("TransferBetweenAccount", accountOriginID, accountDestinationID, amount).Return(accountOriginReturned, nil)

	reqJSON := fmt.Sprintf(`{
				  "account_destination_id": "%v",
				  "amount": %v
				}`, accountDestinationID, amount)

	rec, ctx := testutils.GetRecordedAndContextWithJWT(echo.POST, "/api/transfers", strings.NewReader(reqJSON), accountOriginID)

	handler := Handler{TransferService: serviceMock}
	assert.NoError(t, handler.TransferToAnotherAccount(ctx))
	assert.Equal(t, http.StatusOK, rec.Code)
	t.Log(rec.Body.String())
	m := testutils.ResponseToMap(rec.Body.Bytes())
	assert.Equal(t, accountOriginID, m["id"])
	assert.Equal(t, amount, m["balance"])
}

func Test_TransferToAnotherAccount_ShouldReturnStatusUnauthorized_WhenJWTIsInvalid(t *testing.T) {
	serviceMock := setupService()
	reqJSON := fmt.Sprintf(`{
				  "account_destination_id": "%v",
				  "amount": %v
				}`, accountDestinationID, amount)
	rec, ctx := testutils.GetRecordedAndContext(echo.POST, "/api/transfers", strings.NewReader(reqJSON))
	handler := Handler{TransferService: serviceMock}
	assert.NoError(t, handler.TransferToAnotherAccount(ctx))
	t.Log(rec.Body.String())
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	m := testutils.ResponseToMap(rec.Body.Bytes())
	assert.Equal(t, httputil.HTTPErrorInvalidJWT.Message, m["message"])
}

func Test_TransferToAnotherAccount_ShouldReturnBadRequest_WhenHasErrorOnBody(t *testing.T) {
	serviceMock := setupService()
	reqJSON := fmt.Sprintf(`{
				 	some invalid body
				`)
	rec, ctx := testutils.GetRecordedAndContextWithJWT(echo.POST, "/api/transfers", strings.NewReader(reqJSON), accountOriginID)
	handler := Handler{TransferService: serviceMock}
	assert.NoError(t, handler.TransferToAnotherAccount(ctx))
	t.Log(rec.Body.String())
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	m := testutils.ResponseToMap(rec.Body.Bytes())
	assert.Equal(t, httputil.HTTPErrorParseBody.Message, m["message"])
}

func Test_TransferToAnotherAccount_ShouldReturnBadRequest_WhenBodyMissRequiredFields(t *testing.T) {
	serviceMock := setupService()
	reqJSON := fmt.Sprintf(`{}`)
	rec, ctx := testutils.GetRecordedAndContextWithJWT(echo.POST, "/api/transfers", strings.NewReader(reqJSON), accountOriginID)
	handler := Handler{TransferService: serviceMock}
	assert.NoError(t, handler.TransferToAnotherAccount(ctx))
	t.Log(rec.Body.String())
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	m := testutils.ResponseToMap(rec.Body.Bytes())
	assert.Equal(t, httputil.HTTPErrorValidateBody.Message, m["message"])
}

func Test_TransferToAnotherAccount_ShouldReturnNotFound_WhenGotNotFoundError(t *testing.T) {
	serviceMock := setupService()
	serviceMock.On("TransferBetweenAccount", accountOriginID, accountDestinationID, amount).Return(emptyAccount, errors.Wrap(service.ErrorNotFound, "some error"))
	reqJSON := fmt.Sprintf(`{
				  "account_destination_id": "%v",
				  "amount": %v
				}`, accountDestinationID, amount)
	rec, ctx := testutils.GetRecordedAndContextWithJWT(echo.POST, "/api/transfers", strings.NewReader(reqJSON), accountOriginID)
	handler := Handler{TransferService: serviceMock}
	assert.NoError(t, handler.TransferToAnotherAccount(ctx))
	t.Log(rec.Body.String())
	assert.Equal(t, http.StatusNotFound, rec.Code)
	m := testutils.ResponseToMap(rec.Body.Bytes())
	assert.Contains(t, m["message"], service.ErrorNotFound.Error())
}

func Test_TransferToAnotherAccount_ShouldReturnBadRequest_WhenGotGenericError(t *testing.T) {
	serviceMock := setupService()
	serviceMock.On("TransferBetweenAccount", accountOriginID, accountDestinationID, amount).Return(emptyAccount, errors.New("some error"))
	reqJSON := fmt.Sprintf(`{
				  "account_destination_id": "%v",
				  "amount": %v
				}`, accountDestinationID, amount)
	rec, ctx := testutils.GetRecordedAndContextWithJWT(echo.POST, "/api/transfers", strings.NewReader(reqJSON), accountOriginID)
	handler := Handler{TransferService: serviceMock}
	assert.NoError(t, handler.TransferToAnotherAccount(ctx))
	t.Log(rec.Body.String())
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	m := testutils.ResponseToMap(rec.Body.Bytes())
	assert.Equal(t, "some error", m["message"])
}
