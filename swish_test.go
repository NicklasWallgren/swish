package swish

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"testing"
)

func TestPaymentRequest(t *testing.T) {
	configuration := NewConfiguration(
		&TestEnvironment,
		&Pkcs12{Content: loadPkcs12(getResourcePath("certificates/test.p12")), Password: "qwerty123"},
	)

	payload := PaymentPayload{PayeePaymentReference: "0123456789", CallbackURL: "https://myfakehost.se/swishcallback.cfm", PayeeAlias: "9871065216", PayerAlias: "1231181189", Amount: "100", Currency: "SEK"}

	swish := New(configuration)

	response, err := swish.Payment(context.Background(), &payload)
	if err != nil {
		fmt.Println(err)

		return
	}
	response2, err := swish.PaymentResult(context.Background(), response.ID)
	if err != nil {
		return
	}

	fmt.Println(response2)
}

func getResourcePath(path string) (directory string) {
	dir, err := getResourceDirectoryPath()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/%s", dir, path)
}

func getResourceDirectoryPath() (directory string, err error) {
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		return "", errors.New("could not derive directory path")
	}

	return fmt.Sprintf("%s/%s", path.Dir(filename), "./resource"), nil
}

func loadPkcs12(path string) []byte {
	// #nosec:G304
	cert, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return cert
}
