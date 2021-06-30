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
	return c.Redirect("/history")
}

func (a *App) Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func (a *App) History(c *fiber.Ctx) error {
	histories, err := a.Usecase.WithdrawUC.History(c.UserContext())
	if err != nil {
		return err
	}
	return c.Render("history", fiber.Map{
		"histories": histories,
	})
}
