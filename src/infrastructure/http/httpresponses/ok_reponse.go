package httpresponses

import "net/http"

// NewHTTPOkResponse creates a new created response
func NewHTTPOkResponse(message string, data interface{}) Message {
	return &defaultResponseMessage{
		defaultMessage{message,
			http.StatusOK,
			"created",
		},
		data,
	}
}
