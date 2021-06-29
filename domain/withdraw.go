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

const (
	WithdrawPending = "PENDING"
	WithdrawSuccess = "SUCCESS"
	WithdrawFail    = "FAIL"
)
