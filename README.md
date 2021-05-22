# Swish library

Golang library for the Swish Payment and Refund Request API.

[![Build Status](https://github.com/NicklasWallgren/swish/workflows/Test/badge.svg)](https://github.com/NicklasWallgren/swish/actions?query=workflow%3ATest)
[![Reviewdog](https://github.com/NicklasWallgren/bankid/workflows/reviewdog/badge.svg)](https://github.com/NicklasWallgren/bankid/actions?query=workflow%3Areviewdog)
[![Go Report Card](https://goreportcard.com/badge/github.com/NicklasWallgren/swish)](https://goreportcard.com/report/github.com/NicklasWallgren/swish)
[![GoDoc](https://godoc.org/github.com/NicklasWallgren/swish?status.svg)](https://godoc.org/github.com/NicklasWallgren/swish) 

Check out the API Documentation http://godoc.org/github.com/NicklasWallgren/swish

# Installation
The library can be installed through `go get` 
```bash
go get github.com/NicklasWallgren/swish
```

# Supported versions
We support the two major Go versions, which are 1.14 and 1.15 at the moment.

# Features
- Create payment request
- Retrieve payment result
- Create refund request
- Retrieve refund result

# Examples 

## Initiate payment request
```go
import (
    "context"
    "fmt"
    "io/ioutil"
    "github.com/NicklasWallgren/swish"
)

certificate, err := ioutil.ReadFile("path/to/environment.p12")
if err != nil {
    panic(err)
}

configuration := swish.NewConfiguration(
    &swish.TestEnvironment,
    &swish.Pkcs12{Content: certificate, Password: "p12 password"},
)

instance := swish.New(configuration)

payload := swish.PaymentPayload{PayeePaymentReference: "0123456789", CallbackUrl: "https://myfakehost.se/swishcallback.cfm", PayeeAlias: "9871065216", PayerAlias: "1231181189", Amount: "100", Currency: "SEK"}

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
```

## Unit tests
```bash
go test -v -race $(go list ./... | grep -v vendor)
```

### Code Guide

We use GitHub Actions to make sure the codebase is consistent (`golangci-lint run`) and continuously tested (`go test -v -race $(go list ./... | grep -v vendor)`). We try to keep comments at a maximum of 120 characters of length and code at 120.

## Contributing

If you find any problems or have suggestions about this library, please submit an issue. Moreover, any pull request, code review and feedback are welcome.

## Contributors
  - [Nicklas Wallgren](https://github.com/NicklasWallgren)
  - [All Contributors][link-contributors]

[link-contributors]: ../../contributors

## License

[MIT](./LICENSE)
