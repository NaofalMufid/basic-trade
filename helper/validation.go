package helper

import (
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Msg   string `json:"msg"`
}

func (e ValidationError) Error() string {
	return e.Msg
}

func ConvertValidationErrors(err error) []ValidationError {
	var validationErrors []ValidationError

	if verr, ok := err.(validator.ValidationErrors); ok {
		for _, e := range verr {
			validationErrors = append(validationErrors, ValidationError{
				Field: e.Field(),
				Tag:   e.Tag(),
				Msg:   e.Error(),
			})
		}
	}

	return validationErrors
}
