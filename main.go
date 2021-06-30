package main

import (
	"flip/app"
	"flip/config"
)

func main() {
	// setup config
	cfg := config.NewConfig()

	// create app
	myMarket := app.NewApp(*cfg)
	defer myMarket.DB.Close()

	app.Route(myMarket)
	app.Cron(myMarket)

	// run cron in another goroutine
	myMarket.Cron.Start()

	// run fiber
	myMarket.Fiber.Listen(":8000")
}
