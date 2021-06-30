package app

import (
	"flip/domain"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (a *App) Withdraw(c *fiber.Ctx) error {
	var amount int
	var err error
	if c.FormValue("amount") != "" {
		amount, err = strconv.Atoi(c.FormValue("amount"))
		if err != nil {
			return err
		}
	}
	r := domain.WithdrawRequest{
		AccountNumber: c.FormValue("account_number"),
		BankCode:      c.FormValue("bank_code"),
		Remark:        c.FormValue("remark"),
		Amount:        amount,
	}
	errs := r.Validate()
	if len(errs) > 0 {
		return c.JSON(map[string]interface{}{
			"code": "ERROR_VALIDATION",
			"data": errs,
		})
	}
	_, err = a.Usecase.WithdrawUC.Create(c.UserContext(), r)
	if err != nil {
		return err
	}
	return c.Redirect("/history")
}

func (a *App) Index(c *fiber.Ctx) error {
	return c.Render("template/index", fiber.Map{})
}

func (a *App) History(c *fiber.Ctx) error {
	histories, err := a.Usecase.WithdrawUC.History(c.UserContext())
	if err != nil {
		return err
	}
	return c.Render("template/history", fiber.Map{
		"histories": histories,
	})
}
