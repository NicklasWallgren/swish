package swish

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	location            = "Location"
	paymentRequestToken = "PaymentRequestToken"
)

type Response interface {
	Decode(response *http.Response, swish *Swish) (Response, error)
}

type PaymentResponse struct {
	// Used as reference to this order when the client is started automatically.
	Id string
	// Used to collect the status of the order.
	Url string
	// Payment request token
	Token string
	swish *Swish
}

func (p *PaymentResponse) String() string {
	return fmt.Sprintf("%#v", p)
}

func (p *PaymentResponse) Decode(response *http.Response, swish *Swish) (Response, error) {
	if location := getHttpHeaderValue(location, response); len(location) > 0 {
		id, _ := getIdFromLocation(location)

		p.Id = id
		p.Url = location
	}

	if token := getHttpHeaderValue(paymentRequestToken, response); len(token) > 0 {
		p.Token = token
	}

	return p, nil
}

type PaymentResultResponse struct {
	Id                    string  `json:"id"`                    // id
	PayeePaymentReference string  `json:"payeePaymentReference"` // payeePaymentReference
	PaymentReference      string  `json:"paymentReference"`      // paymentReference
	CallbackUrl           string  `json:"callbackUrl"`           // callbackUrl
	PayerAlias            string  `json:"payerAlias"`            // payerAlias
	PayeeAlias            string  `json:"payeeAlias"`            // payeeAlias
	Amount                float32 `json:"amount"`                // amount
	Currency              string  `json:"currency"`              // currency
	Message               string  `json:"message"`               // message
	Status                string  `json:"status"`                // status
	DateCreated           string  `json:"dateCreated"`           // dateCreated
	DatePaid              string  `json:"datePaid"`              // datePaid
	ErrorCode             string  `json:"errorCode"`             // errorCode
	ErrorMessage          string  `json:"errorMessage"`          // errorMessage
}

func (p *PaymentResultResponse) String() string {
	return fmt.Sprintf("%#v", p)
}

func (p *PaymentResultResponse) Decode(response *http.Response, swish *Swish) (Response, error) {
	decoder := json.NewDecoder(response.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&p)

	return p, err
}

type RefundResponse struct {
}

func (RefundResponse) String() string {
	panic("implement me")
}

func (RefundResponse) Decode(response *http.Response, swish *Swish) (Response, error) {
	panic("implement me")
}

type RefundResultResponse struct {
}

func (RefundResultResponse) String() string {
	panic("implement me")
}

func (RefundResultResponse) Decode(response *http.Response, swish *Swish) (Response, error) {
	panic("implement me")
}

type ErrorResponse []Error

type Error struct {
	Code                  string `json:"ErrorCode""`
	Message               string `json:"ErrorMessage"`
	AdditionalInformation string `json:"additionalInformation"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("%#v", e)
}

func (e *ErrorResponse) Decode(response *http.Response, swish *Swish) (Response, error) {
	decoder := json.NewDecoder(response.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&e)

	return e, err
}
