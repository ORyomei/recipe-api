package controllers

import "net/http"

// Message has a response message
type Message struct {
	Message string `json:"message"`
}

// Error is error for api
type Error struct {
	Message string `json:"message"`
	Type    string `json:"type,omitempty"`
	Status  int    `json:"-"`
}

// NewBadRequestError create BadRequestError
func NewBadRequestError(message string) *Error {
	return &Error{
		Message: message,
		Status:  http.StatusBadRequest,
	}
}

// NewNotFoundError create NotFoundError
func NewNotFoundError(message string) *Error {
	return &Error{
		Message: message,
		Status:  http.StatusNotFound,
	}
}

// NewInternalServerError create InternalServerError
func NewInternalServerError() *Error {
	return &Error{
		Message: "エラーが発生しました",
		Status:  http.StatusInternalServerError,
	}
}

// NewUnauthorizedError create Unauthorized
func NewUnauthorizedError() *Error {
	return &Error{
		Message: "認証に失敗しました",
		Status:  http.StatusUnauthorized,
	}
}

type APIErrorType string

const (
	ErrorTypeStationClosed   = APIErrorType("station_closed")
	ErrorTypeAlreadyAcquired = APIErrorType("already_acquired")
)

// NewDomainError create Error about domain
func NewDomainError(t APIErrorType) *Error {
	return &Error{
		Status: http.StatusBadRequest,
		Type:   string(t),
	}
}
