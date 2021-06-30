package domain

import "time"

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
