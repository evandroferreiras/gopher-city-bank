package api

import (
	"github.com/evandroferreiras/gopher-city-bank/app/modules/account"
	"github.com/evandroferreiras/gopher-city-bank/app/modules/healthcheck"
	"github.com/evandroferreiras/gopher-city-bank/app/modules/login"
)

// Handler is the struct for API
type Handler struct {
	healthCheck *healthcheck.Handler
	account     *account.Handler
	login       *login.Handler
}

// NewHandler is a constructor to API Handler
func NewHandler() *Handler {
	return &Handler{
		healthCheck: healthcheck.NewHandler(),
		account:     account.NewHandler(),
		login:       login.NewHandler(),
	}
}
