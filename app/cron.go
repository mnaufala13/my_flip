package app

import (
	"context"
	"log"
)

func Cron(app *App) {
	app.Cron.AddFunc("CRON_TZ=Asia/Jakarta * * * * *", func() {
		log.Println("running sync withdrawal")
		err := app.Usecase.WithdrawUC.SyncWithdrawal(context.Background())
		if err != nil {
			log.Println(err)
		}
	})
	app.Cron.AddFunc("CRON_TZ=Asia/Jakarta * * * * *", func() {
		log.Println("running sync bigflip transaction")
		err := app.Usecase.BigflipUC.Sync(context.Background())
		if err != nil {
			log.Println(err)
		}
	})
}
