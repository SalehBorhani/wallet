package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func FormatValidationError(err error) []ErrorMsg {
<<<<<<< HEAD
	errorMessages := []ErrorMsg{}

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		e := strings.Split(err.Error(), ".")[1]
		errors := strings.Split(e, " of type ")
		message := fmt.Sprintf("%s must be %s", errors[0], errors[1])
		errorMessages = append(errorMessages, ErrorMsg{Field: errors[0], Message: message})
=======
	var errorMessages []ErrorMsg

	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		e := strings.Split(err.Error(), ".")[1]
		split := strings.Split(e, " of type ")
		message := fmt.Sprintf("%s must be %s", split[0], split[1])
		errorMessages = append(errorMessages, ErrorMsg{Field: split[0], Message: message})
>>>>>>> a0a5309 (fix(): add atomic transaction to the wallet)
		return errorMessages
	}

	if !errors.Is(err, io.EOF) {
		for _, e := range err.(validator.ValidationErrors) {
			var message string
			if e.Param() != "" {
				message = ToSnakeCase(e.Field()) + " " + e.Tag() + " " + e.Param()
			} else {
				message = ToSnakeCase(e.Field()) + " " + e.Tag()
			}
			errorMessages = append(errorMessages, ErrorMsg{Field: ToSnakeCase(e.Field()), Message: message})
		}
		return errorMessages
	}
	return errorMessages
}
