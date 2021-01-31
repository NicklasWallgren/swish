package swish

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	location            = "Location"
	paymentRequestToken = "PaymentRequestToken"
)

// Response is the interface implemented by types that holds the response context fields.
type Response interface {
	Decode(response *http.Response, swish *Swish) (Response, error)
}

// PaymentResponse holds the information of a initiated payment request.
type PaymentResponse struct {
	// Used as reference to this order when the client is started automatically.
	ID string `json:"id"`
	// Used to collect the status of the order.
	URL string `json:"url"`
	// Payment request token
	Token string
	// nolint:structcheck,unused
	swish *Swish
}

func (p *PaymentResponse) String() string {
	return fmt.Sprintf("%#v", p)
}

// Decode reads the http response and stories it in a payment response struct.
func (p *PaymentResponse) Decode(response *http.Response, swish *Swish) (Response, error) {
	if location := getHTTPHeaderValue(location, response); len(location) > 0 {
		id, _ := getIDFromLocation(location)

		p.ID = id
		p.URL = location
	}

	if token := getHTTPHeaderValue(paymentRequestToken, response); len(token) > 0 {
		p.Token = token
	}

	return p, nil
}

// PaymentResultResponse holds the information of a initiated payment result request.
type PaymentResultResponse struct {
	ID                    string  `json:"id"`
	PayeePaymentReference string  `json:"payeePaymentReference"`
	PaymentReference      string  `json:"paymentReference"`
	CallbackURL           string  `json:"callbackUrl"`
	PayerAlias            string  `json:"payerAlias"`
	PayeeAlias            string  `json:"payeeAlias"`
	Amount                float32 `json:"amount"`
	Currency              string  `json:"currency"`
	Message               string  `json:"message"`
	Status                string  `json:"status"`
	DateCreated           string  `json:"dateCreated"`
	DatePaid              string  `json:"datePaid"`
	ErrorCode             string  `json:"errorCode"`
	ErrorMessage          string  `json:"errorMessage"`
}

func (p *PaymentResultResponse) String() string {
	return fmt.Sprintf("%#v", p)
}

// Decode reads the JSON-encoded response and stories it in a payment result response struct.
func (p *PaymentResultResponse) Decode(response *http.Response, swish *Swish) (Response, error) {
	return p, decode(response.Body, p)
}

// RefundResponse contains fields specific for the refund response.
type RefundResponse struct {
	// Used as reference to this order when the client is started automatically.
	ID string `json:"id"`
	// Used to collect the status of the refund.
	URL string `json:"url"`
	// nolint:structcheck,unused
	swish *Swish
}

func (r RefundResponse) String() string {
	return fmt.Sprintf("%#v", r)
}

// Decode reads the JSON-encoded value and stories it in a refund response struct.
func (r *RefundResponse) Decode(response *http.Response, swish *Swish) (Response, error) {
	if location := getHTTPHeaderValue(location, response); len(location) > 0 {
		id, _ := getIDFromLocation(location)

		r.ID = id
		r.URL = location
	}

	return r, nil
}

// RefundResultResponse contains fields specific for the refund result response.
type RefundResultResponse struct {
	ID                       string  `json:"id"`
	PaymentReference         string  `json:"paymentReference"`
	PayerPaymentReference    string  `json:"payerPaymentReference"`
	OriginalPaymentReference string  `json:"originalPaymentReference"`
	CallbackURL              string  `json:"callbackUrl"`
	PayerAlias               string  `json:"payerAlias"`
	PayeeAlias               string  `json:"payeeAlias"`
	Amount                   float32 `json:"amount"`
	Currency                 string  `json:"currency"`
	Message                  string  `json:"message"`
	Status                   string  `json:"status"`
	DateCreated              string  `json:"dateCreated"`
	DatePaid                 string  `json:"datePaid"`
	ErrorMessage             string  `json:"errorMessage"`
	AdditionalInformation    string  `json:"additionalInformation"`
	ErrorCode                string  `json:"errorCode"`
}

func (r RefundResultResponse) String() string {
	return fmt.Sprintf("%#v", r)
}

// Decode reads the JSON-encoded value and stories it in a refund result response struct.
func (r *RefundResultResponse) Decode(response *http.Response, swish *Swish) (Response, error) {
	return r, decode(response.Body, r)
}

// ErrorResponse contains fields for the error response.
type ErrorResponse []Error

// Error contains fields specific for the API error.
type Error struct {
	Code                  string `json:"ErrorCode"`
	Message               string `json:"ErrorMessage"`
	AdditionalInformation string `json:"additionalInformation"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("%#v", e)
}

// Decode reads the JSON-encoded value and stories it in a error response struct.
func (e *ErrorResponse) Decode(response *http.Response, swish *Swish) (Response, error) {
	return e, decode(response.Body, e)
}

func decode(subject io.ReadCloser, target interface{}) error {
	decoder := json.NewDecoder(subject)
	decoder.DisallowUnknownFields()

	return decoder.Decode(&target)
}
