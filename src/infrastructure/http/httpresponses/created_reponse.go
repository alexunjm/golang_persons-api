package httpresponses

import "net/http"

// NewHTTPCreatedResponse creates a new created response
func NewHTTPCreatedResponse(message string) Message {
	return &defaultMessage{
		message,
		http.StatusCreated,
		"created",
	}
}
