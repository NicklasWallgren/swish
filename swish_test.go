package swish

import (
	"context"
	"fmt"
	"testing"
)

func TestPaymentRequest(t *testing.T) {
	configuration := NewConfiguration(&TestEnvironment, GetResourcePath("certificates/test.pem"), GetResourcePath("certificates/test.key"))
	payload := PaymentPayload{PayeePaymentReference: "0123456789", CallbackUrl: "https://myfakehost.se/swishcallback.cfm", PayeeAlias: "9871065216", PayerAlias: "1231181189", Amount: "100", Currency: "SEK"}

	swish := New(configuration)

	response, err := swish.Payment(context.Background(), &payload)

	if err != nil {
		fmt.Println(err)

		return
	}
	response2, err := swish.PaymentResult(context.Background(), response.Id)

	if err != nil {
		return
	}

	fmt.Println(response2)

}
