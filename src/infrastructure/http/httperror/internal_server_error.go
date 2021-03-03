package exception

import "net/http"

// NewInternalServerError creates a new internal server error
func NewInternalServerError(message string) MessageErr {
	return &messageErr{
		ErrMessage:    message,
		ErrHTTPStatus: http.StatusInternalServerError,
		ErrCode:       "server_error",
	}
}
