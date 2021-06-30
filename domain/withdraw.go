package domain

import (
	"time"
)

type WithdrawModel struct {
	Id            string
	AccountNumber string
	BankCode      string
	Remark        string
	Amount        int
	CreatedAt     time.Time
}

type WithdrawRequest struct {
	AccountNumber string
	BankCode      string
	Remark        string
	Amount        int
}

func (r WithdrawRequest) Validate() []string {
	errs := []string{}
	if r.AccountNumber == "" {
		errs = append(errs, "account number can't empty")
	}
	if r.BankCode == "" {
		errs = append(errs, "bank code can't empty")
	}
	if r.Remark == "" {
		errs = append(errs, "remark can't empty")
	}
	if r.Amount == 0 {
		errs = append(errs, "amount can't zero")
	}
	return errs
}

type History struct {
	Amount        int    `json:"amount"`
	Status        string `json:"status"`
	BankCode      string `json:"bank_code"`
	AccountNumber string `json:"account_number"`
	Remark        string `json:"remark"`
	Receipt       string `json:"receipt"`
	TimeServed    string `json:"time_served"`
	Fee           int    `json:"fee"`
}

const (
	WithdrawPending = "PENDING"
	WithdrawSuccess = "SUCCESS"
	WithdrawFail    = "FAIL"
)
