package cmd

import (
	"context"
	"flip/app"
	"flip/config"
	"flip/domain"
	"github.com/spf13/cobra"
	"log"
)

var withdrawCmd = &cobra.Command{
	Use:   "withdraw",
	Short: "Run Service",
	Long:  `Run Service`,
	Run: func(cmd *cobra.Command, args []string) {
		// setup config
		cfg := config.NewConfig()

		// create app
		myMarket := app.NewApp(*cfg, templates)
		defer myMarket.DB.Close()

		_, err := myMarket.Usecase.WithdrawUC.Create(context.Background(), domain.WithdrawRequest{
			AccountNumber: accountNumber,
			BankCode:      bankCode,
			Remark:        remark,
			Amount:        int(amount),
		})
		if err != nil {
			log.Println(err)
		}
	},
}

var (
	accountNumber, bankCode, remark string
	amount                          int64
)

func init() {
	withdrawCmd.Flags().StringVarP(&accountNumber, "account_number", "c", "", "account number value")
	withdrawCmd.Flags().StringVarP(&bankCode, "bank_code", "b", "", "bank code value")
	withdrawCmd.Flags().StringVarP(&remark, "remark", "r", "", "remark value")
	withdrawCmd.Flags().Int64VarP(&amount, "amount", "a", 0, "amount value")
	rootCmd.AddCommand(withdrawCmd)
}
