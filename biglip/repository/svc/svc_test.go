package svc

import (
	"context"
	"flip/domain"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func Test_buildRequest(t *testing.T) {
	request, err := buildRequest("GET", baseUrl+"/users", nil)
	assert.NoError(t, err)
	assert.Equal(t, "GET", request.Method)
	assert.Equal(t, baseUrl, fmt.Sprintf("%s://%s", request.URL.Scheme, request.Host))
	assert.Equal(t, "/users", request.URL.Path)
	assert.Equal(t, "application/x-www-form-urlencoded", request.Header.Get("Content-Type"))
}

func TestFlipper_callDisburse(t *testing.T) {
	c := &http.Client{}
	flipper := flipper{client: c}
	payload := domain.DisbursePayload{
		BankCode:      "bni",
		AccountNumber: "110011212",
		Amount:        1000,
		Remark:        "this is remark",
	}
	disburse, err := flipper.callDisburse(payload)
	assert.NoError(t, err)
	assert.EqualValues(t, payload.BankCode, disburse.BankCode)
	assert.EqualValues(t, payload.AccountNumber, disburse.AccountNumber)
	assert.EqualValues(t, payload.Amount, disburse.Amount)
	assert.EqualValues(t, payload.Remark, disburse.Remark)
}

func Test_flipper_Status(t *testing.T) {
	c := &http.Client{}
	flipper := flipper{client: c}
	id := 5535152564
	flipTrx, err := flipper.Status(context.Background(), id)
	assert.NoError(t, err)
	assert.EqualValues(t, id, flipTrx.Id)
}

func Test_flipTransactionToModel(t *testing.T) {
	ft := domain.FlipTransaction{
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
	tModel, err := flipTransactionToModel(ft)
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
	dt, err := normalizeDatetime(nullDt)
	assert.NoError(t, err)
	assert.EqualValues(t, "0001-01-01T00:00:00Z", dt)
}
