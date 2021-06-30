package domain

import (
	"flip/models"
	"fmt"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"strings"
	"time"
)

type FlipTransaction struct {
	Id              int64  `json:"id"`
	Amount          int    `json:"amount"`
	Status          string `json:"status"`
	Timestamp       string `json:"timestamp"`
	BankCode        string `json:"bank_code"`
	AccountNumber   string `json:"account_number"`
	BeneficiaryName string `json:"beneficiary_name"`
	Remark          string `json:"remark"`
	Receipt         string `json:"receipt"`
	TimeServed      string `json:"time_served"`
	Fee             int    `json:"fee"`
}

type DisbursePayload struct {
	BankCode      string
	AccountNumber string
	Amount        int
	Remark        string
}

func NormalizeDatetime(datetime string) (string, error) {
	if datetime == "0000-00-00 00:00:00" {
		return "0001-01-01T00:00:00Z", nil
	}
	tt := strings.Split(datetime, " ")
	if len(tt) != 2 {
		return "", errors.New("invalid format")
	}
	dt := fmt.Sprintf("%sT%s+07:00", tt[0], tt[1])
	return dt, nil
}

func FlipTransactionToModel(ft FlipTransaction) (*models.BigflipLog, error) {
	tn, err := NormalizeDatetime(ft.TimeServed)
	if err != nil {
		return nil, err
	}
	t, err := time.Parse(time.RFC3339, tn)
	if err != nil {
		return nil, errors.Wrap(err, "error parse time serve")
	}
	ttn, err := NormalizeDatetime(ft.Timestamp)
	if err != nil {
		return nil, err
	}
	trxTimestamp, err := time.Parse(time.RFC3339, ttn)
	if err != nil {
		return nil, errors.Wrap(err, "error parse transaction timestamp")
	}
	lg := &models.BigflipLog{
		TransactionID:   ft.Id,
		Amount:          ft.Amount,
		Status:          ft.Status,
		TRXTimestamp:    null.TimeFrom(trxTimestamp),
		BankCode:        ft.BankCode,
		AccountNumber:   ft.AccountNumber,
		BeneficiaryName: ft.BeneficiaryName,
		Remark:          ft.Remark,
		Receipt:         ft.Receipt,
		TimeServed:      null.TimeFrom(t),
		Fee:             ft.Fee,
	}
	return lg, nil
}
