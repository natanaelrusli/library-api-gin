package apperror

import (
	"errors"
	"fmt"
)

func NewNotFoundError(err error, entityName string) AppError {
	msg := fmt.Sprintf("%s not found", entityName)

	if err == nil {
		err = errors.New(msg)
	}

	return NewAppError(err, NotFoundErrorCode, msg, nil)
}
