package main

import (
	"context"
	"fmt"
	"github.com/NicklasWallgren/swish"
)

func main() {
	configuration := swish.NewConfiguration(
		&swish.TestEnvironment,
		swish.GetResourcePath("certificates/test.pem"),
		swish.GetResourcePath("certificates/test.key"))

	instance := swish.New(configuration)

	refundPayload := swish.RefundPayload{
		OriginalPaymentReference: "ID",
		CallbackUrl:              "https://myfakehost.se/swishcallback.cfm",
		PayerAlias:               "9871065216",
		PayeeAlias:               "1231181189",
		Amount:                   "100",
		Currency:                 "SEK",
		Message:                  "Refund",
	}

	// Initiates a refund request.
	refundResponse, err := instance.Refund(context.Background(), &refundPayload)

	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(refundResponse)

	// Retrieves the refund result
	refundResultResponse, err := instance.RefundResult(context.Background(), refundResponse.Id)

	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(refundResultResponse)
}
