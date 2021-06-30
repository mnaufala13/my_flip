package app

func Route(app *App) {
	f := app.Fiber
	f.Get("/", app.Index)
	f.Get("/history", app.History)
	f.Post("/withdraw", app.Withdraw)
}
