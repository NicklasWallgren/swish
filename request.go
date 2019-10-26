package swish

const (
	uriPayment = "paymentrequests"
	uriRefund  = "refunds"
)

type Request interface {
	Method() string
	Uri() string
	Payload() payloadInterface
	Response() Response
}

type request struct {
	method   string
	uri      string
	payload  payloadInterface
	response Response
}

func (r request) Method() string {
	return r.method
}

func (r request) Uri() string {
	return r.uri
}

func (r request) Payload() payloadInterface {
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
