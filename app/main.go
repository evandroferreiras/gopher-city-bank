package main

import (
	"github.com/evandroferreiras/gopher-city-bank/app/api"
	"github.com/evandroferreiras/gopher-city-bank/app/common/log"
)

// @title Gopher City Bank API
// @version 1.0

// @contact.name Evandro souza
// @contact.email evandroferreiras@gmail.com

func main() {
	log.Init()

	echo := api.New()
	v1 := echo.Group("/api")
	h := api.NewHandler()
	h.Register(v1)
	h.RegisterSwagger(echo)
	echo.Logger.Fatal(echo.Start(":8585"))
}
