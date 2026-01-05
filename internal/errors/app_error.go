package errors

import "net/http"

type AppError struct {
	StatusCode int
	Code       string
	Message    string
	err        error
}

func (e *AppError) Error() string {
	return e.Message
}

func New(status int, code, message string, err error) *AppError {
	return &AppError{
		StatusCode: status,
		Code:       code,
		Message:    message,
		err:        err,
	}
}

// Common helpers
func BadRequest(message string, err error) *AppError {
	return New(http.StatusBadRequest, "BAD_REQUEST", message, err)
}

func NotFound(message string, err error) *AppError {
	return New(http.StatusNotFound, "NOT_FOUND", message, err)
}

func Internal(message string, err error) *AppError {
	return New(http.StatusInternalServerError, "INTERNAL_ERROR", message, err)
}
func Unauthorized(message string, err error) *AppError {
	return New(http.StatusUnauthorized, "UNAUTHORIZED", message, err)
}

func Forbidden(message string, err error) *AppError {
	return New(http.StatusForbidden, "FORBIDDEN", message, err)
}
func Conflict(message string, err error) *AppError {
	return New(http.StatusConflict, "CONFLICT", message, err)
}
func TooManyRequests(message string, err error) *AppError {
	return New(http.StatusTooManyRequests, "TOO_MANY_REQUESTS", message, err)
}
