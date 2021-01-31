package swish

import "encoding/json"

type encoder interface {
	encode(payload payload) ([]byte, error)
}

type jsonEncoder struct{}

func newJSONEncoder() encoder {
	return &jsonEncoder{}
}

func (e jsonEncoder) encode(payload payload) ([]byte, error) {
	return json.Marshal(payload)
}
