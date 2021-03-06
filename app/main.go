package main

import (
	"github.com/evandroferreiras/gopher-city-bank/app/api"
	"github.com/evandroferreiras/gopher-city-bank/app/api/server"
	"github.com/evandroferreiras/gopher-city-bank/app/common/envvar"
	"github.com/evandroferreiras/gopher-city-bank/app/common/log"
	"github.com/evandroferreiras/gopher-city-bank/app/db"
	"github.com/sirupsen/logrus"
)

// @title Gopher City Bank API
// @version 1.0
// @contact.name Evandro Souza
// @contact.email evandroferreiras@gmail.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log.Init()

	if envvar.ExecuteAutoMigrate() {
		db.AutoMigrate()
		logrus.Info("AUTOMIGRATE executed")
		return
	}

	echo := server.New()
	v1 := echo.Group("/api")
	h := api.NewHandler()
	h.Register(v1)
	h.RegisterSwagger(echo)
	echo.Logger.Fatal(echo.Start(":8585"))
}
