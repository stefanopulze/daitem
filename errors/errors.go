package errors

import "fmt"

type ApiError struct {
	Code    int
	Message string
	Cause   error
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func GenericError(code int, message string) error {
	return &ApiError{
		Code:    code,
		Message: message,
	}
}

func HttpError(err error) error {
	return &ApiError{
		Code:    HTTP,
		Cause:   err,
		Message: err.Error(),
	}
}

func JsonError(err error) error {
	return &ApiError{
		Code:    JSON,
		Cause:   err,
		Message: err.Error(),
	}
}
