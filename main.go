package main

import (
	"flip/app"
	"flip/config"
)

func main() {
	cfg := config.NewConfig()
	myMarket := app.NewApp(*cfg)
	app.Route(myMarket)
	myMarket.Fiber.Listen(":8000")
	defer myMarket.DB.Close()
}
