package svc

import (
	"context"
	"flip/domain"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_buildRequest(t *testing.T) {
	request, err := buildRequest("GET", baseUrl+"/users", nil)
	assert.NoError(t, err)
	assert.Equal(t, "GET", request.Method)
	assert.Equal(t, baseUrl, fmt.Sprintf("%s://%s", request.URL.Scheme, request.Host))
	assert.Equal(t, "/users", request.URL.Path)
	assert.Equal(t, "application/x-www-form-urlencoded", request.Header.Get("Content-Type"))
}

func TestFlipper_Disburse(t *testing.T) {
	c := &http.Client{}
	flipper := NewFlipper(c)
	payload := domain.DisbursePayload{
		BankCode:      "bni",
		AccountNumber: "110011212",
		Amount:        1000,
		Remark:        "this is remark",
	}
	disburse, err := flipper.Disburse(context.Background(), "", payload)
	assert.NoError(t, err)
	assert.EqualValues(t, payload.BankCode, disburse.BankCode)
	assert.EqualValues(t, payload.AccountNumber, disburse.AccountNumber)
	assert.EqualValues(t, payload.Amount, disburse.Amount)
	assert.EqualValues(t, payload.Remark, disburse.Remark)
}

func Test_flipper_Status(t *testing.T) {
	c := &http.Client{}
	flipper := NewFlipper(c)
	id := 5535152564
	flipTrx, err := flipper.Status(context.Background(), id)
	assert.NoError(t, err)
	assert.EqualValues(t, id, flipTrx.Id)
}
