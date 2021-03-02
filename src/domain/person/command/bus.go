package command

// Bus is a contract to handle commands through a bus
type Bus interface {
	Register(commandType Type, handler Handler)
	// getHandler(commandType Type) Handler
	Dispatch(command Command)
}
