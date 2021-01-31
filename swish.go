package swish

import (
	"context"
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

// Swish contains the validator and configuration context.
type Swish struct {
	validator     *validator.Validate
	configuration *Configuration
	client        Client
}

// New returns a new instance of 'Swish'.
func New(configuration *Configuration) Swish {
	return Swish{validator: newValidator(), configuration: configuration}
}

// Payment - Initiates a payment request
//
// Merchants and Technical Suppliers can send payment requests for both E-Commerce and M- Commerce to MSS.
//
// Once MSS receives a “payment request” call, there are two answers that will be returned from MSS (unless error situation).
// The first answer is synchronous, the second one is asynchronous.
func (s Swish) Payment(ctx context.Context, payload *PaymentPayload) (*PaymentResponse, error) {
	request := newPaymentRequest(payload)

	response, err := s.call(ctx, request)
	if err != nil {
		return nil, err
	}

	paymentResponse := (response).(*PaymentResponse)
	return paymentResponse, nil
}

// PaymentResult - Retrieves the payment result
//
// The client can retrieve payment result information of an initiated payment request (successful or failed)
// If the previous create payment request call simulated a delayed error the response of the GET operation will
// have a status property with value ERROR and properties errorCode and errorMessage will be set accordingly
//
// MSS stores the necessary information about each incoming “PaymentRequest request” in a cache which automatically expires
// every 24 hours or when the MSS server is restarted.
func (s Swish) PaymentResult(ctx context.Context, token string) (*PaymentResultResponse, error) {
	request := newPaymentResultRequest(token)

	response, err := s.call(ctx, request)
	if err != nil {
		return nil, err
	}

	paymentResultResponse := (response).(*PaymentResultResponse)
	return paymentResultResponse, nil
}

// Refund - Initiates a refund request.
//
// Merchants and Technical Suppliers can send refund request to MSS.
//
// When MSS receives a “Refund request” there are three answers that will be returned from MSS (unless error situation).
// The first answer is synchronous, the second and third responses are asynchronous.
func (s Swish) Refund(ctx context.Context, payload *RefundPayload) (*RefundResponse, error) {
	request := newRefundRequest(payload)

	response, err := s.call(ctx, request)
	if err != nil {
		return nil, err
	}

	refundResponse := (response).(*RefundResponse)
	return refundResponse, nil
}

// RefundResult - Retrieves the refund result
//
// The client can retrieve refund result information of an initiated refund request (successful or failed).
//
// If the previous create refund request call simulated a delayed error the response of the GET operation will
// have a status property with value ERROR and properties errorCode and errorMessage will be set accordingly
//
// MSS stores the necessary information about each incoming “Refund request” in a cache which automatically expires
// every 24 hours or when the MSS server is restarted.
func (s Swish) RefundResult(ctx context.Context, token string) (*RefundResultResponse, error) {
	request := newRefundResultRequest(token)

	response, err := s.call(ctx, request)
	if err != nil {
		return nil, err
	}

	refundResultResponse := (response).(*RefundResultResponse)
	return refundResultResponse, nil
}

func (s Swish) call(ctx context.Context, request Request) (Response, error) {
	if err := s.validate(request); err != nil {
		return nil, nil
	}

	if err := s.initialize(); err != nil {
		return nil, err
	}

	return (s.client).call(ctx, request, &s)
}

func (s Swish) validate(request Request) error {
	if request.Payload() == nil {
		return nil
	}

	// Validate the integrity of the payload
	if err := s.validator.Struct(request.Payload()); err != nil {
		return fmt.Errorf("payload validation error %w", err)
	}

	return nil
}

// initialize prepares the client in head of a request.
func (s *Swish) initialize() error {
	// Check whether the client has been initialized
	if s.client != nil {
		return nil
	}

	// Lazy initialization
	client, err := newClient(s.configuration)
	if err != nil {
		return err
	}

	s.client = client

	return nil
}
