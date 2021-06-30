package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_flipTransactionToModel(t *testing.T) {
	ft := FlipTransaction{
		Id:              161877137,
		Amount:          11,
		Status:          "PENDING",
		Timestamp:       "2021-06-30 15:16:44",
		BankCode:        "bni",
		AccountNumber:   "10102020",
		BeneficiaryName: "PT FLIP",
		Remark:          "eeee",
		Receipt:         "",
		TimeServed:      "0000-00-00 00:00:00",
		Fee:             400,
	}
	tModel, err := FlipTransactionToModel(ft)
	assert.NoError(t, err)
	assert.EqualValues(t, ft.Id, tModel.TransactionID)
	assert.EqualValues(t, ft.Amount, tModel.Amount)
	assert.EqualValues(t, ft.Status, tModel.Status)
	assert.EqualValues(t, "2021-06-30T15:16:44+07:00", tModel.TRXTimestamp.Time.Format(time.RFC3339))
	assert.EqualValues(t, ft.BankCode, tModel.BankCode)
	assert.EqualValues(t, ft.AccountNumber, tModel.AccountNumber)
	assert.EqualValues(t, ft.BeneficiaryName, tModel.BeneficiaryName)
	assert.EqualValues(t, ft.Remark, tModel.Remark)
	assert.EqualValues(t, ft.Receipt, tModel.Receipt)
	assert.EqualValues(t, "0001-01-01T00:00:00Z", tModel.TimeServed.Time.Format(time.RFC3339))
	assert.EqualValues(t, ft.Fee, tModel.Fee)
}

func Test_normalizeDatetime(t *testing.T) {
	nullDt := "0000-00-00 00:00:00"
	dt, err := NormalizeDatetime(nullDt)
	assert.NoError(t, err)
	assert.EqualValues(t, "0001-01-01T00:00:00Z", dt)
}
