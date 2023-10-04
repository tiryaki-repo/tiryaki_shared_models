package errors

import (
	"fmt"
	"net/http"
)

var (
	ErrBadRequest = HttpError{
		Status:  Status(http.StatusBadRequest),
		Message: http.StatusText(http.StatusBadRequest),
		Err:     fmt.Errorf(http.StatusText(http.StatusBadRequest))}

	ErrInternalServer = HttpError{
		Status:  Status(http.StatusInternalServerError),
		Message: http.StatusText(http.StatusInternalServerError),
		Err:     fmt.Errorf(http.StatusText(http.StatusInternalServerError))}

	ErrNotFound = HttpError{
		Status:  Status(http.StatusNotFound),
		Message: http.StatusText(http.StatusNotFound),
		Err:     fmt.Errorf(http.StatusText(http.StatusNotFound))}
)

type Status int

type HttpError struct {
	Status  Status
	Message string
	Err     error
}

func (se HttpError) Error() string {
	return se.Message
}

func (se HttpError) Unwrap() error {
	return se.Err
}
