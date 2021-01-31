package swish

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Client is the interface implemented by types that can invoke the BankID REST API.
type Client interface {
	// call is responsible for making the HTTP call against BankId REST API
	call(context context.Context, request Request, swish *Swish) (Response, error)
}

type client struct {
	client        *http.Client
	configuration *Configuration
	encoder       encoder
	decoder       decoder
}

func (c client) call(context context.Context, request Request, swish *Swish) (Response, error) {
	encoded, err := c.encoder.encode(request.Payload())
	if err != nil {
		return nil, fmt.Errorf("unable to encode payload %w", err)
	}

	req, err := c.newRequest(context, request.Method(), c.urlFrom(request), strings.NewReader(string(encoded)))
	if err != nil {
		return nil, err
	}

	resp, err := c.request(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close() // nolint:errcheck

	return c.decoder.decode(request.Response(), resp, swish)
}

// ClientOption definition.
type ClientOption func(*client)

// newClient returns a new instance of 'newClient'.
func newClient(configuration *Configuration, options ...ClientOption) (Client, error) {
	clientCfg, err := newTLSClientConfig(configuration)
	if err != nil {
		return nil, fmt.Errorf("error reading and/or parsing the certification files. Cause: %w", err)
	}

	netClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: clientCfg,
		},
	}

	instance := &client{client: &netClient, configuration: configuration, encoder: newJSONEncoder(), decoder: newJSONDecoder()}

	// Apply options if there are any, can overwrite default
	for _, option := range options {
		option(instance)
	}

	return instance, nil
}

// Function to create ConfigurationOption func to set net/http client.
// nolint:deadcode, unused
func withHTTPClient(target *http.Client) ClientOption {
	return func(subject *client) {
		subject.client = target
	}
}

// newRequest creates and prepares a instance of http request.
func (c client) newRequest(ctx context.Context, method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		return nil, fmt.Errorf("unable to build request %w", err)
	}

	return req, nil
}

func (c client) urlFrom(request Request) string {
	return c.configuration.Environment.BaseURL + "/" + request.URI()
}

func (c client) request(request *http.Request) (*http.Response, error) {
	return c.client.Do(request)
}

func isHTTPStatusCodeWithinRange(statusCode int, statusCodeRange statusCodeRange) bool {
	return statusCode >= statusCodeRange.start && statusCode <= statusCodeRange.end
}

func getHTTPHeaderValue(header string, response *http.Response) string {
	return response.Header.Get(header)
}
