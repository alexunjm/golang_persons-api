package exception

import "net/http"

// NewUnprocessableEntityError creates a new unprocessable entity error
func NewUnprocessableEntityError(message string) MessageErr {
	return &messageErr{
		ErrMessage:    message,
		ErrHTTPStatus: http.StatusUnprocessableEntity,
		ErrCode:       "invalid_request",
	}
}
