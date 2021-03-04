package httpresponses

import "net/http"

// NewHTTPNoContentResponse creates a new created response
func NewHTTPNoContentResponse(message string) Message {
	return &defaultMessage{
		message,
		http.StatusNoContent,
		"no_content",
	}
}
