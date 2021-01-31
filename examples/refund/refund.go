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

	refundPayload := swish.RefundPayload{
		OriginalPaymentReference: "ID",
		CallbackURL:              "https://myfakehost.se/swishcallback.cfm",
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
	refundResultResponse, err := instance.RefundResult(context.Background(), refundResponse.ID)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(refundResultResponse)
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
