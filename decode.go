package swish

import (
	"fmt"
	"net/http"
)

var successRange = statusCodeRange{200, 300 - 1}

type decoder interface {
	decode(subject Response, response *http.Response, swish *Swish) (Response, error)
}

type jsonDecoder struct{}

func newJSONDecoder() decoder {
	return &jsonDecoder{}
}

func (j jsonDecoder) decode(subject Response, response *http.Response, swish *Swish) (Response, error) {
	if isHTTPStatusCodeWithinRange(response.StatusCode, successRange) {
		decoded, err := subject.Decode(response, swish)
		if err != nil {
			return nil, fmt.Errorf("unable to decode response %w", err)
		}

		return decoded, nil
	}

	return nil, j.decodeError(response, swish)
}

func (j jsonDecoder) decodeError(response *http.Response, swish *Swish) error {
	errorResponse := ErrorResponse{}

	_, err := errorResponse.Decode(response, swish)
	if err != nil {
		return fmt.Errorf("%w. Error code %d. Response: %s", err, response.StatusCode, tryReadCloserToString(response.Body))
	}

	return &errorResponse
}

type statusCodeRange struct {
	start int
	end   int
}
