package httpresponses

// Message interface for standard responses
type Message interface {
	Text() string
	Status() int
	StringCode() string
}
