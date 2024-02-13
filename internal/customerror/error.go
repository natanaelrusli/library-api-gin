package customerror

import (
	"fmt"
)

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (ce *CustomError) Error() string {
	return fmt.Sprintf("error %d: %s", ce.Code, ce.Message)
}

func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}
