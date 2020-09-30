package server

import (
	"github.com/evandroferreiras/gopher-city-bank/app/common/log"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	logMiddleware "github.com/neko-neko/echo-logrus/v2"
)

// New builds a new instance of echo
func New() *echo.Echo {
	e := echo.New()
	e.Pre(echoMiddleware.RemoveTrailingSlash())
	e.Logger = log.EchoLogrus()
	e.Use(logMiddleware.Logger())

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Validator = NewValidator()
	return e
}
