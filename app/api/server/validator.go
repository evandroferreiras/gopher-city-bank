package server

import (
	"gopkg.in/go-playground/validator.v9"
)

// NewValidator is a constructor for Validator
func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// Validator is a struct for validate based on structs
type Validator struct {
	validator *validator.Validate
}

// Validate validates a structs exposed fields, and automatically validates nested structs, unless otherwise specified
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
