package command

type Handler interface {
	Handle(command Command) error
}
