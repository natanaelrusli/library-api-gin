package apperror

import "errors"

func NewBookNotFoundError() AppError {
	msg := "book not found"

	err := errors.New(msg)

	return NewAppError(err, DefaultClientErrorCode, msg, nil)
}
