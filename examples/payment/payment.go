package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"

	"github.com/NicklasWallgren/swish"
)

func main() {
	configuration := swish.NewConfiguration(
		&swish.TestEnvironment,
		&swish.Pkcs12{Content: loadPkcs12(getResourcePath("certificates/test.p12")), Password: "swish"},
	)

	instance := swish.New(configuration)

	payload := swish.PaymentPayload{
		PayeePaymentReference: "0123456789",
		CallbackURL:           "https://myfakehost.se/swishcallback.cfm",
		PayeeAlias:            "9871065216",
		PayerAlias:            "1231181189",
		Amount:                "100",
		Currency:              "SEK",
	}

	// Initiates a payment request
	paymentResponse, err := instance.Payment(context.Background(), &payload)
	if err != nil {
		fmt.Println(err)

		return
	}

	// Retrieves the payment result
	paymentResult, err := instance.PaymentResult(context.Background(), paymentResponse.ID)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(paymentResult)
}

func loadPkcs12(path string) []byte {
	cert, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return cert
}

// getResourceDirectoryPath returns the full path to the resource directory.
func getResourceDirectoryPath() (directory string, err error) {
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		return "", fmt.Errorf("could not derive directory path")
	}

	return fmt.Sprintf("%s/%s", path.Dir(filename), "../../resource"), nil
}

// getResourcePath returns the full path to the resource.
func getResourcePath(path string) (directory string) {
	dir, err := getResourceDirectoryPath()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/%s", dir, path)
}
