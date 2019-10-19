# Swish library

Golang library for the Swish Payment and Refund Request API.

[![Build Status](https://travis-ci.org/NicklasWallgren/swish.svg?branch=master)](https://travis-ci.org/NicklasWallgren/swish)
[![Go Report Card](https://goreportcard.com/badge/github.com/stretchr/testify)](https://goreportcard.com/report/github.com/NicklasWallgren/swish)
[![GoDoc](https://godoc.org/github.com/NicklasWallgren/swish?status.svg)](https://godoc.org/github.com/NicklasWallgren/swish) 

Check out the API Documentation http://godoc.org/github.com/NicklasWallgren/swish

# Installation
The library can be installed through `go get` 
```bash
go get github.com/NicklasWallgren/swish
```

# Supported versions
We support the two major Go versions, which are 1.12 and 1.13 at the moment.

# Features
- Create payment request
- Retrieve payment result
- Create refund request
- Retrieve refund result

# Examples 

## Initiate payment request
```go
import (
    "fmt"
    "github.com/NicklasWallgren/swish"
)

configuration := swish.NewConfiguration(
    &swish.TestEnvironment,
    swish.GetResourcePath("certificates/test.pem"),
    swish.GetResourcePath("certificates/test.key"))

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

## TODO
 - Add unit tests
 - Add validator translator

## Contributing
  - Fork it!
  - Create your feature branch: `git checkout -b my-new-feature`
  - Commit your changes: `git commit -am 'Useful information about your new features'`
  - Push to the branch: `git push origin my-new-feature`
  - Submit a pull request

## Contributors
  - [Nicklas Wallgren](https://github.com/NicklasWallgren)
  - [All Contributors][link-contributors]

[link-contributors]: ../../contributors