package app

func Route(app *App) {
	f := app.Fiber
	f.Post("/withdraw", app.Withdraw)
}
