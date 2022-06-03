package apierrors

import "net/http"

type APIError interface {
	Error() string
	Status() int
}

type apiError struct {
	ErrorMessage string `json:"error_message"`
	StatusCode   int    `json:"status_code"`
}

func (apiErr apiError) Error() string {
	return apiErr.ErrorMessage
}

func (apiErr apiError) Status() int {
	return apiErr.StatusCode
}

func NewBadRequestAPIError(msg string) APIError {
	return apiError{
		ErrorMessage: msg,
		StatusCode:   http.StatusBadRequest,
	}
}

func NewUnauthorizedAPIError(msg string) APIError {
	return apiError{
		ErrorMessage: msg,
		StatusCode:   http.StatusUnauthorized,
	}
}

func NewNotFoundAPIError(msg string) APIError {
	return apiError{
		ErrorMessage: msg,
		StatusCode:   http.StatusNotFound,
	}
}

func NewInternalServerAPIError(msg string) APIError {
	return apiError{
		ErrorMessage: msg,
		StatusCode:   http.StatusInternalServerError,
	}
}
