package exception

// MessageErr interface for standard exception responses
type MessageErr interface {
	Text() string
	Status() int
	Error() string
}
