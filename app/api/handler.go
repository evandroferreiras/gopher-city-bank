package api

import "github.com/evandroferreiras/gopher-city-bank/app/healthcheck"

// Handler is the struct for API
type Handler struct {
	healthCheck *healthcheck.Handler
}

// NewHandler is a constructor to API Handler
func NewHandler() *Handler {
	return &Handler{
		healthCheck: healthcheck.NewHandler(),
	}
}
