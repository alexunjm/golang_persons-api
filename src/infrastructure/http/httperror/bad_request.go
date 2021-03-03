package exception

import "net/http"

// NewBadRequestError creates a new bad request error
func NewBadRequestError(message string) MessageErr {
	return &messageErr{
		ErrMessage:    message,
		ErrHTTPStatus: http.StatusBadRequest,
		ErrCode:       "bad_request",
	}
}
