package main

import (
	"context"
	"fmt"
	"swish"
)

func main() {
	configuration := swish.NewConfiguration(
		&swish.TestEnvironment,
		swish.GetResourcePath("certificates/test.pem"),
		swish.GetResourcePath("certificates/test.key"))

	instance := swish.New(configuration)

	payload := swish.PaymentPayload{
		PayeePaymentReference: "0123456789",
		CallbackUrl:           "https://myfakehost.se/swishcallback.cfm",
		PayeeAlias:            "9871065216",
		PayerAlias:            "1231181189",
		Amount:                "100",
		Currency:              "SEK"}

	paymentResponse, err := instance.Payment(context.Background(), &payload)

	if err != nil {
		fmt.Println(err)

		return
	}

	paymentResult, err := instance.PaymentResult(context.Background(), paymentResponse.Id)

	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(paymentResult)
}
