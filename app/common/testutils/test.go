package testutils

import (
	"encoding/json"
	"io"
	"net/http/httptest"

	myEcho "github.com/evandroferreiras/gopher-city-bank/app/api/server"
	"github.com/labstack/echo/v4"
)

// GetRecordedAndContext returns response recorder and context to help unit tests
func GetRecordedAndContext(method string, target string, body io.Reader) (*httptest.ResponseRecorder, echo.Context) {
	e := myEcho.New()
	req := httptest.NewRequest(method, target, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return rec, c
}

// ResponseToMap converts a Json response body to Map. Helps to test the returned value
func ResponseToMap(b []byte) map[string]interface{} {
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	return m
}
