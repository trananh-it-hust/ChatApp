package response

import (
	"net/http"
)

type AppError struct {
	Code      int    `json:"code"`
	RootError error  `json:"-"`
	Message   string `json:"message"`
	Log       string `json:"detail"`
	Key       string `json:"key"`
}

func NewAppError(rootError error, message string, log string, key string) *AppError {
	return &AppError{
		Code:      http.StatusBadRequest,
		RootError: rootError,
		Message:   message,
		Log:       log,
		Key:       key,
	}
}

func NewAppErrorWithCode(code int, rootError error, message string, log string, key string) *AppError {
	return &AppError{
		Code:      code,
		RootError: rootError,
		Message:   message,
		Log:       log,
		Key:       key,
	}
}
func (e *AppError) GetRootError() error {
	if err, ok := e.RootError.(*AppError); ok {
		return err.GetRootError()
	}
	return e.RootError
}

func (e *AppError) Error() string {
	return e.RootError.Error()
}

func ErrorNotFound(err error) *AppError {
	return NewAppErrorWithCode(http.StatusNotFound, err, "Not Found", err.Error(), "not_found")
}
func ErrorBadRequest(err error) *AppError {
	return NewAppErrorWithCode(http.StatusBadRequest, err, "Bad Request", err.Error(), "bad_request")
}
func ErrorInternalServer(err error) *AppError {
	return NewAppErrorWithCode(http.StatusInternalServerError, err, "Internal Server Error", err.Error(), "internal_server")
}
func ErrorUnauthorized(err error) *AppError {
	return NewAppErrorWithCode(http.StatusUnauthorized, err, "Unauthorized", err.Error(), "unauthorized")
}
func ErrorForbidden(err error) *AppError {
	return NewAppErrorWithCode(http.StatusForbidden, err, "Forbidden", err.Error(), "forbidden")
}
