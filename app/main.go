package main

import (
	"github.com/evandroferreiras/gopher-city-bank/app/api"
	"github.com/evandroferreiras/gopher-city-bank/app/common/log"
)

func main() {
	log.Init()

	echo := api.New()
	v1 := echo.Group("/api")
	h := api.NewHandler()
	h.Register(v1)
	echo.Logger.Fatal(echo.Start("127.0.0.1:8585"))
}
