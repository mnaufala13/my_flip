package app

import (
	"flip/domain"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (a *App) Withdraw(c *fiber.Ctx) error {
	amount, err := strconv.Atoi(c.FormValue("amount"))
	if err != nil {
		return err
	}
	_, err = a.Usecase.WithdrawUC.Create(c.UserContext(), domain.WithdrawRequest{
		AccountNumber: c.FormValue("account_number"),
		BankCode:      c.FormValue("bank_code"),
		Remark:        c.FormValue("remark"),
		Amount:        amount,
	})
	if err != nil {
		return err
	}
	return nil
}
