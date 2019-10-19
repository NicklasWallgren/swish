package swish

import (
	"golang.org/x/xerrors"
	"net/http"
)

var (
	successRange = statusCodeRange{200, 300 - 1}
	errorRange   = statusCodeRange{400, 600 - 1}
)

type Decoder interface {
	decode(subject Response, response *http.Response, swish *Swish) (*Response, error)
}

type jsonDecoder struct{}

func newJsonDecoder() Decoder {
	return &jsonDecoder{}
}

func (j jsonDecoder) decode(subject Response, response *http.Response, swish *Swish) (*Response, error) {
	if isHttpStatusCodeWithinRange(response.StatusCode, successRange) {
		decoded, err := subject.Decode(response, swish)

		return &decoded, err
	}

	if isHttpStatusCodeWithinRange(response.StatusCode, errorRange) {
		return nil, j.decodeError(response, swish)
	}

	decoded, err := subject.Decode(response, swish)

	return &decoded, err
}

func (j jsonDecoder) decodeError(response *http.Response, swish *Swish) error {
	errorResponse := ErrorResponse{}

	_, err := errorResponse.Decode(response, swish)

	if err != nil {
		return xerrors.Errorf("%w. Error code %d. Response: %s", err, response.StatusCode, response.Body)
	}

	return &errorResponse
}

type statusCodeRange struct {
	start int
	end   int
}
