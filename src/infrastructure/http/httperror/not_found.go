package exception

import "net/http"

// NewNotFoundError creates a new not found error
func NewNotFoundError(message string) MessageErr {
	return &messageErr{
		ErrMessage:    message,
		ErrHTTPStatus: http.StatusNotFound,
		ErrCode:       "not_found",
	}
}
