package cmd

import (
	"embed"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "appmain",
	Short: "Test bigflip by Naufal",
	Long:  `Test bigflip by Naufal`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func Execute(files embed.FS) {
	templates = files
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

}
