package command

type Bus interface {
	Register(commandType Type, handler Handler) error
	getHandler(commandType Type) *Handler
	Dispatch(command Command) error
}
