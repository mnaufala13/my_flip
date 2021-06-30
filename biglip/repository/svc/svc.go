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
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
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

func NewFlipper(client *http.Client, db *sql.DB, psqlFlip postgres.BigFlipPsqlRepository) BigFlipSvcRepository {
	return &flipper{client: client, db: db, psqlFlip: psqlFlip}
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
	flipTrx, err := f.callDisburse(payload)
	if err != nil {
		return nil, err
	}
	go func() {
		_, err := f.log(ctx, withdrawalId, *flipTrx)
		if err != nil {
			log.Println(err)
		}
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

func normalizeDatetime(datetime string) (string, error) {
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

func flipTransactionToModel(ft domain.FlipTransaction) (*models.BigflipLog, error) {
	tn, err := normalizeDatetime(ft.TimeServed)
	if err != nil {
		return nil, err
	}
	t, err := time.Parse(time.RFC3339, tn)
	if err != nil {
		return nil, errors.Wrap(err, "error parse time serve")
	}
	ttn, err := normalizeDatetime(ft.Timestamp)
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

func (f *flipper) log(ctx context.Context, withdrawalId string, ft domain.FlipTransaction) (*models.BigflipLog, error) {
	lg, err := flipTransactionToModel(ft)
	if err != nil {
		return nil, err
	}
	lg.WithdrawalID = withdrawalId
	_, err = f.psqlFlip.Insert(ctx, f.db, *lg)
	if err != nil {
		return nil, err
	}
	return lg, nil
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
