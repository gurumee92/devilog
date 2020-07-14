package router

import "github.com/go-playground/validator"

// NewValidator is create validator instance
func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// Validator is validate dto
type Validator struct {
	validator *validator.Validate
}

// Validate is validate model
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
