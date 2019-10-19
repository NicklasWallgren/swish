package swish

// payloadInterface is the interface implemented by types that holds the fields to be delivered to the API
type payloadInterface interface{}

// payload holds the request fields to be delivered to the API
type payload struct {
	payloadInterface
}

type PaymentPayload struct {
	*payload
	// Payment reference supplied by theMerchant. This is not used by Swish but is included in responses back to the
	// client. This reference could for example be an order id or similar. If set the value must not exceed 35 characters
	// and only the following characters are allowed: [a-ö, A-Ö, 0-9, -]
	PayeePaymentReference string `json:"payeePaymentReference,omitempty"` // validate
	// URL that Swish will use to notify caller about the result of the payment request. The URL has to use HTTPS.
	CallbackUrl string `json:"callbackUrl"` // validate
	// The registered Cell phone number of the person that makes the payment. It can only contain numbers and has to be
	// at least 8 and at most 15 digits. It also needs to match the following format in order to be found in
	// Swish: country code + cell phone number (without leading zero). E.g.: 46712345678
	// If set, request is handled as E-Commerce payment.
	// If not set, request is handled as M- Commerce payment.
	PayerAlias string `json:"payerAlias,omitempty"` // validate
	// The social security number of the individual making the payment,
	// should match the registered value for payerAlias or the payment will not be accepted.
	PayerSSN string `json:"payerSSN,omitempty"`
	// Minimum age (in years) that the individual connected to the payerAlias has to be in order for the payment to
	// be accepted. Value has to be in the range of 1 to 99.
	AgeLimit string `json:"ageLimit,omitempty"` // validate
	// The Swish number of the payee. It needs to match with Merchant Swish number.
	PayeeAlias string `json:"payeeAlias"` // validate
	// The amount of money to pay. The amount cannot be less than 1 SEK and not more than
	// 999999999999.99 SEK. Valid value has to
	// be all digits or with 2 digit decimal separated with a period.
	Amount string `json:"amount"` // validate
	// The currency to use. Currently the only supported value is SEK.
	Currency string `json:"currency"` // validate
	// Merchant supplied message about the payment/order. Max 50 characters.
	// Allowed characters are the letters a-ö, A-Ö, the numbers 0-9 and any of the special characters :;.,?!()-”.
	Message string `json:"message,omitempty"` // validate`
}

type RefundPayload struct {
	*payload
}
