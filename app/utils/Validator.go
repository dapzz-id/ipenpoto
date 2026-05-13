package utils

import (
	"github.com/go-playground/validator/v10"
)

var ValidateRequest = validator.New()

func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {
		errors[e.Field()] = e.Tag()
	}

	return errors
}