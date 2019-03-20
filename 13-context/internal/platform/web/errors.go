package web

import (
	"net/http"

	"github.com/pkg/errors"
)

// errorResponse is the form used for API responses from failures in the API.
type errorResponse struct {
	Error string `json:"error"`
}

// statusError is used to pass errors through the application with web specific
// context.
type statusError struct {
	err    error
	status int
}

// ErrorWithStatus wraps a provided error with an HTTP status code.
func ErrorWithStatus(err error, status int) error {
	return &statusError{err, status}
}

// Error implements the error interface. It uses the default message of the
// wrapped error. This is what will be shown in the services' logs.
func (se *statusError) Error() string {
	return se.err.Error()
}

// ExternalError provides "human readable" error messages that are intended for
// service users to see. If the status code is 500 or higher (the default) then
// a generic error message is returned.
//
// The idea is that a developer who creates an error like this intends to let
// the API consumer know the product was not found:
//	ErrorWithStatus(errors.New("product not found"), 404)
//
// However a more serious error like a database failure might include
// information that is not safe to show to API consumers.
func (se *statusError) ExternalError() string {
	if se.status < http.StatusInternalServerError {
		return se.err.Error()
	}
	return http.StatusText(se.status)
}

// toStatusError takes a regular error and converts it to a statusError. If the
// original error is already a *statusError it is returned directly. If not
// then it is defaulted to an error with a 500 status.
func toStatusError(err error) *statusError {
	if se, ok := errors.Cause(err).(*statusError); ok {
		return se
	}
	return &statusError{err, http.StatusInternalServerError}
}
