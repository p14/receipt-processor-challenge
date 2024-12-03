package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Validator wraps the go-playground validator instance.
type Validator struct {
	Validate *validator.Validate
}

// NewValidator initializes and returns a Validator.
func NewValidator() *Validator {
	v := validator.New()

	// Register custom validation for currency (up to two decimal places)
	v.RegisterValidation("currency", currencyValidationFunc)

	return &Validator{
		Validate: v,
	}
}

// currencyValidationFunc ensures that the string represents a number with up to two decimal places.
func currencyValidationFunc(fl validator.FieldLevel) bool {
	// Define a regex pattern for numbers with up to two decimal places
	pattern := `^\d+(\.\d{1,2})?$`
	matched, _ := regexp.MatchString(pattern, fl.Field().String())
	return matched
}
