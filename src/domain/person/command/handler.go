package command

// Handler is a contract to handle commands
type Handler interface {
	Handle(command Command) error
}
