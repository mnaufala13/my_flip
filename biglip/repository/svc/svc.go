package svc

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"flip/biglip/repository/postgres"
	"flip/domain"
	"flip/models"
	"fmt"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type BigFlipSvcRepository interface {
	Disburse(ctx context.Context, withdrawalId string, payload domain.DisbursePayload) (*domain.FlipTransaction, error)
	Status(ctx context.Context, transactionId int) (*domain.FlipTransaction, error)
}

const baseUrl = "https://nextar.flip.id"

type flipper struct {
	client   *http.Client
	db       *sql.DB
	psqlFlip postgres.BigFlipPsqlRepository
}

func NewFlipper(c *http.Client) BigFlipSvcRepository {
	return &flipper{client: c}
}

func buildRequest(method, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth("HyzioY7LP6ZoO7nTYKbG8O4ISkyWnX1JvAEVAhtWKZumooCzqp41", "")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return request, nil
}

// Disburse call the bigflip api and then log the response
func (f *flipper) Disburse(ctx context.Context, withdrawalId string, payload domain.DisbursePayload) (*domain.FlipTransaction, error) {
	tx, err := f.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error begin trx at disburse")
	}
	flipTrx, err := f.callDisburse(payload)
	if err != nil {
		return nil, err
	}
	go func() {
		_, err := f.log(ctx, tx, withdrawalId, *flipTrx)
		log.Println(err)
	}()
	return flipTrx, nil
}

// Status to get status bigflip transaction
func (f *flipper) Status(ctx context.Context, transactionId int) (*domain.FlipTransaction, error) {
	const endpoint = baseUrl + "/disburse"

	flipTrx := &domain.FlipTransaction{}

	request, err := buildRequest("GET", fmt.Sprintf("%s/%d", endpoint, transactionId), nil)
	if err != nil {
		return nil, err
	}

	response, err := f.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(flipTrx)
	if err != nil {
		return nil, err
	}

	return flipTrx, nil
}

func (f *flipper) log(ctx context.Context, exec boil.ContextExecutor, withdrawalId string, ft domain.FlipTransaction) (*models.BigflipLog, error) {
	lg := models.BigflipLog{
		TransactionID:   ft.Id,
		Amount:          ft.Amount,
		Status:          ft.Status,
		TRXTimestamp:    null.TimeFrom(ft.Timestamp),
		BankCode:        ft.BankCode,
		AccountNumber:   ft.AccountNumber,
		BeneficiaryName: ft.BeneficiaryName,
		Remark:          ft.Remark,
		Receipt:         ft.Receipt,
		TimeServed:      null.TimeFrom(ft.TimeServed),
		Fee:             ft.Fee,
		WithdrawalID:    withdrawalId,
	}
	_, err := f.psqlFlip.Insert(ctx, exec, lg)
	if err != nil {
		return nil, err
	}
	return &lg, nil
}

func (f flipper) callDisburse(payload domain.DisbursePayload) (*domain.FlipTransaction, error) {
	const endpoint = baseUrl + "/disburse"

	disburse := &domain.FlipTransaction{}

	var param = url.Values{}
	param.Set("bank_code", payload.BankCode)
	param.Set("remark", payload.Remark)
	param.Set("account_number", payload.AccountNumber)
	param.Set("amount", strconv.Itoa(payload.Amount))

	var flipPayload = bytes.NewBufferString(param.Encode())
	request, err := buildRequest("POST", endpoint, flipPayload)
	if err != nil {
		return nil, err
	}

	response, err := f.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		bodyString := string(bodyBytes)
		return nil, errors.New(fmt.Sprintf("error disburse: %s", bodyString))
	}

	err = json.NewDecoder(response.Body).Decode(disburse)
	if err != nil {
		return nil, err
	}

	return disburse, nil
}
