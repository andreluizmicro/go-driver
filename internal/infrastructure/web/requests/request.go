package request

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type RequestError struct {
	Error error
	Tag   string
	Field string
}

func ValidateRequest(w http.ResponseWriter, input any) validator.ValidationErrors {
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return err.(validator.ValidationErrors)
		// for _, e := range errors {

		// 	setError(w, request.RequestError(e.ActualTag(), e.Field()), http.StatusUnprocessableEntity)
		// }
		// return
	}
	return nil
}

func ValidationErrors(errorType, field string) error {
	switch errorType {
	case "required":
		return fmt.Errorf("%s is required", field)
	case "min":
		return fmt.Errorf("field %s must have at least 3 characters", field)
	case "max":
		return fmt.Errorf("field %s must have a maximum of 100 characters", field)
	default:
		return errors.New("internal server error")
	}
}
