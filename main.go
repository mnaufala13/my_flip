package main

import (
	"embed"
	"flip/app"
	"flip/config"
)

//go:embed template/index.html template/history.html
var templates embed.FS

func main() {
	// setup config
	cfg := config.NewConfig()

	// create app
	myMarket := app.NewApp(*cfg, templates)
	defer myMarket.DB.Close()

	app.Route(myMarket)
	app.Cron(myMarket)

	// run cron in another goroutine
	myMarket.Cron.Start()

	// run fiber
	myMarket.Fiber.Listen(":8000")
}
