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
	"github.com/google/uuid"
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

type flipper struct {
	client   *http.Client
	db       *sql.DB
	psqlFlip postgres.BigFlipPsqlRepository
	baseUrl  string
	Secret   string
}

func NewFlipper(client *http.Client, db *sql.DB, psqlFlip postgres.BigFlipPsqlRepository, baseUrl string, secret string) BigFlipSvcRepository {
	return &flipper{client: client, db: db, psqlFlip: psqlFlip, baseUrl: baseUrl, Secret: secret}
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
	flipTrx, err := f.callDisburse(withdrawalId, payload)
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
func (f *flipper) Status(_ context.Context, transactionId int) (*domain.FlipTransaction, error) {
	var endpoint = f.baseUrl + "/disburse"

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

func (f *flipper) log(ctx context.Context, withdrawalId string, ft domain.FlipTransaction) (*models.BigflipLog, error) {
	lg, err := domain.FlipTransactionToModel(ft)
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

func (f flipper) callDisburse(withdrawalId string, payload domain.DisbursePayload) (*domain.FlipTransaction, error) {
	var endpoint = f.baseUrl + "/disburse"

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

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)

	// log every response
	go func() {
		lg := models.BigflipResponse{
			ID:           uuid.NewString(),
			URL:          endpoint,
			Payload:      null.JSONFrom(bodyBytes),
			WithdrawalID: withdrawalId,
			HTTPCode:     response.StatusCode,
		}
		err := lg.Insert(context.Background(), f.db, boil.Infer())
		if err != nil {
			log.Println(err, "error log response bigflip")
		}
	}()

	if response.StatusCode >= 400 {
		return nil, errors.New(fmt.Sprintf("error disburse: %s", bodyString))
	}

	err = json.Unmarshal(bodyBytes, disburse)
	if err != nil {
		return nil, errors.Wrap(err, "error decode disburse response")
	}

	return disburse, nil
}
