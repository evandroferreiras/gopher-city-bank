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

// IsOnline godoc
// @Summary IsOnline
// @Description Returns true or false, depeding on the state of app.
// @Success 200 {integer} string "WORKING"
// @Failure 400 {integer} string "NOT WORKING"
// @Router /api/healthcheck [get]
// IsOnline returns if the app is WORKING or NOT WORKING
func (h *Handler) IsOnline(c echo.Context) error {
	if h.healthCheckService.IsWorking() {
		return c.String(http.StatusOK, "WORKING")
	}
	return c.String(http.StatusBadRequest, "NOT WORKING")
}
