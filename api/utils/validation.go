package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

type ValidationError struct {
	Field   string      `json:"field"`
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
}

func ListErrors(w http.ResponseWriter, err error) {
	var errors []ValidationError

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		panic("Invalid Validation Error")
	}

	for _, err := range validationErrors {
		errors = append(errors, ValidationError{
			Field:   err.Field(),
			Message: getErrorMessage(err.Tag(), err.Field(), err.Param()),
			Value:   err.Value(),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(errors)
}

func getErrorMessage(tag string, field string, param string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%v is required", field)
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("%v must have atleast %v characters", field, param)
	case "max":
		return fmt.Sprintf("%v can't exceed %v characters", field, param)
	default:
		return "Invalid value"
	}
}
