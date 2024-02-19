package apperror

import "errors"

const (
	DefaultServerErrorCode = iota
	DefaultClientErrorCode
	ValidationErrorCode
	NotFoundErrorCode
)

type AppError struct {
	err  error
	code int
	msg  string
	data map[string]interface{}
}

func NewAppError(err error, code int, msg string, data map[string]interface{}) AppError {
	return AppError{
		err:  err,
		code: code,
		msg:  msg,
		data: data,
	}
}

func (e AppError) Error() string {
	if e.msg == "" {
		return "error occured"
	}
	return e.msg
}

func (e AppError) GetCode() int {
	return e.code
}

func (e AppError) OriginalError() error {
	var currErr AppError

	currErr = e

	for {
		nextErr := currErr.err
		if nextErr == nil {
			break
		}

		var appErr AppError
		if !errors.As(nextErr, &appErr) {
			return nextErr
		}
		currErr = appErr
	}

	return e
}

func (e AppError) OriginalMessage() string {
	return e.OriginalError().Error()
}

func (e AppError) DisplayMessage() string {
	return e.msg
}

func (e AppError) AditionalData() map[string]interface{} {
	return e.data
}
