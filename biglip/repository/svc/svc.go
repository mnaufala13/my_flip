package svc

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flip/domain"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type DisbursePayload struct {
	BankCode      string
	AccountNumber string
	Amount        int
	Remark        string
}

type BigFlipRepository interface {
	Disburse(ctx context.Context, payload DisbursePayload) (*domain.FlipDisburse, error)
	Status(ctx context.Context, transactionId int) (*domain.FlipTransaction, error)
}

const baseUrl = "https://nextar.flip.id"

type flipper struct {
	client *http.Client
}

func NewFlipper(c *http.Client) *flipper {
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

func (f *flipper) Disburse(ctx context.Context, payload DisbursePayload) (*domain.FlipDisburse, error) {
	const endpoint = baseUrl + "/disburse"

	disburse := &domain.FlipDisburse{}

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
