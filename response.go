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

type Response interface {
	Decode(response *http.Response, swish *Swish) (Response, error)
}

// PaymentResponse holds the information of a initiated payment request
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

// Decode reads the http response and stories it in a payment response struct
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

// PaymentResponse holds the information of a initiated payment result request
type PaymentResultResponse struct {
	Id                    string  `json:"id"`
	PayeePaymentReference string  `json:"payeePaymentReference"`
	PaymentReference      string  `json:"paymentReference"`
	CallbackUrl           string  `json:"callbackUrl"`
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

// Decode reads the JSON-encoded response and stories it in a payment result response struct
func (p *PaymentResultResponse) Decode(response *http.Response, swish *Swish) (Response, error) {
	return p, decode(response.Body, p)
}

type RefundResponse struct {
	// Used as reference to this order when the client is started automatically.
	Id string
	// Used to collect the status of the refund.
	Url   string
	swish *Swish
}

func (r RefundResponse) String() string {
	return fmt.Sprintf("%#v", r)
}

func (r *RefundResponse) Decode(response *http.Response, swish *Swish) (Response, error) {
	if location := getHttpHeaderValue(location, response); len(location) > 0 {
		id, _ := getIdFromLocation(location)

		r.Id = id
		r.Url = location
	}

	return r, nil
}

type RefundResultResponse struct {

}

func (r RefundResultResponse) String() string {
	return fmt.Sprintf("%#v", r)
}

func (r *RefundResultResponse) Decode(response *http.Response, swish *Swish) (Response, error) {
	return r, decode(response.Body, r)
}

type ErrorResponse []Error

type Error struct {
	Code                  string `json:"ErrorCode"`
	Message               string `json:"ErrorMessage"`
	AdditionalInformation string `json:"additionalInformation"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("%#v", e)
}

// Decode reads the JSON-encoded value and stories it in a error response struct
func (e *ErrorResponse) Decode(response *http.Response, swish *Swish) (Response, error) {
	return e, decode(response.Body, e)
}

func decode(subject io.ReadCloser, target interface{}) error {
	decoder := json.NewDecoder(subject)
	decoder.DisallowUnknownFields()

	return decoder.Decode(&target)
}
