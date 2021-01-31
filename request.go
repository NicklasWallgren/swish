package swish

const (
	uriPayment = "paymentrequests"
	uriRefund  = "refunds"
)

// Request is the interface implemented by types that holds the request context fields.
type Request interface {
	Method() string
	URI() string
	Payload() payload
	Response() Response
}

type request struct {
	method   string
	uri      string
	payload  payload
	response Response
}

func (r request) Method() string {
	return r.method
}

func (r request) URI() string {
	return r.uri
}

func (r request) Payload() payload {
	return r.payload
}

func (r request) Response() Response {
	return r.response
}

func newPaymentRequest(payload *PaymentPayload) Request {
	return &request{method: "POST", uri: uriPayment, payload: payload, response: &PaymentResponse{}}
}

func newPaymentResultRequest(token string) Request {
	return &request{method: "GET", uri: uriPayment + "/" + token, response: &PaymentResultResponse{}}
}

func newRefundRequest(payload *RefundPayload) Request {
	return &request{method: "POST", uri: uriRefund, payload: payload, response: &RefundResponse{}}
}

func newRefundResultRequest(token string) Request {
	return &request{method: "GET", uri: uriRefund + "/" + token, response: &RefundResultResponse{}}
}
