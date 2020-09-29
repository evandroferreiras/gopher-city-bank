package api

import "github.com/labstack/echo/v4"

// Register sets the routes configuration for echo
func (h *Handler) Register(v1 *echo.Group) {
	healthCheck := v1.Group("/healthcheck")
	healthCheck.GET("", h.healthCheck.IsOnline)
}
