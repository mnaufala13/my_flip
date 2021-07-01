package cmd

import (
	"embed"
	"flip/app"
	"flip/config"
	"github.com/spf13/cobra"
)

var templates embed.FS

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Run Service",
	Long:  `Run Service`,
	Run: func(cmd *cobra.Command, args []string) {

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
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)
}
