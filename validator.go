package swish

import (
	"gopkg.in/go-playground/validator.v9"
)

// newValidator returns a new instance of 'Validator' prepared with custom validator methods.
func newValidator() *validator.Validate {
	instance := validator.New()

	return instance
}
