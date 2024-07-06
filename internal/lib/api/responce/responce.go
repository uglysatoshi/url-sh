package responce

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Responce struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK() Responce {
	return Responce{
		Status: StatusOK,
	}
}

func Error(msg string) Responce {
	return Responce{
		Status: StatusError,
		Error:  msg,
	}
}

func ValidationError(errs validator.ValidationErrors) Responce {
	var errMessages []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMessages = append(errMessages, fmt.Sprintf("field %s is a required field", err.Field()))
		case "url":
			errMessages = append(errMessages, fmt.Sprintf("field %s is not valid URL", err.Field()))
		default:
			errMessages = append(errMessages, fmt.Sprintf("field %s is not valid", err.Field()))
		}
	}

	return Responce{
		Status: StatusError,
		Error:  strings.Join(errMessages, ", "),
	}
}
