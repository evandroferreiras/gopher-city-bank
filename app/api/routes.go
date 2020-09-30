package api

import (
	// Register docs from swagger
	_ "github.com/evandroferreiras/gopher-city-bank/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Register sets the routes configuration for echo
func (h *Handler) Register(v1 *echo.Group) {
	healthCheck := v1.Group("/healthcheck")
	healthCheck.GET("", h.healthCheck.IsOnline)

	accounts := v1.Group("/accounts")
	accounts.POST("", h.account.CreateAccount)
}

// RegisterSwagger sets the route to swagger documentations
func (h *Handler) RegisterSwagger(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
