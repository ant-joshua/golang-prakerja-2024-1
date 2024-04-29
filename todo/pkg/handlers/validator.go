package handlers

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

// NewValidator create new validator
func NewValidator() *validator.Validate {

	Validate = validator.New(validator.WithRequiredStructEnabled())

	return Validate
}
