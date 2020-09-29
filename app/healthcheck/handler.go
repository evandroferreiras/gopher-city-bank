package healthcheck

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler is a struct to healthcheck handler
type Handler struct {
	healthCheckService Service
}

// NewHandler is a constructor to HealthCheck Handler
func NewHandler() *Handler {
	return &Handler{
		healthCheckService: NewService(),
	}
}

// IsOnline returns if the app is WORKING or NOT WORKING
func (h *Handler) IsOnline(c echo.Context) error {
	if h.healthCheckService.IsWorking() {
		return c.String(http.StatusOK, "WORKING")
	}
	return c.String(http.StatusOK, "NOT WORKING")
}
