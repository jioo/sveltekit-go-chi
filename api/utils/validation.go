package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrorList struct {
	Errors []ValidationError `json:"errors"`
}

func ListErrors(w http.ResponseWriter, err error) {
	var errors []ValidationError
	var errorList ValidationErrorList

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		panic("Invalid Validation Error")
	}

	for _, err := range validationErrors {
		errors = append(errors, ValidationError{
			Field:   err.Field(),
			Message: getErrorMessage(err.Tag(), err.Field(), err.Param()),
		})
	}
	errorList.Errors = errors

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(errorList)
}

func CustomError(w http.ResponseWriter, msg string) {
	var errors []ValidationError
	var errorList ValidationErrorList

	errors = append(errors, ValidationError{
		Field:   "General",
		Message: msg,
	})
	errorList.Errors = errors

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(errorList)
}

func getErrorMessage(tag string, field string, param string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%v is required", field)
	case "min":
		return fmt.Sprintf("%v must have atleast %v characters", field, param)
	default:
		return "Invalid value"
	}
}
