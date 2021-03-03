package command

import "context"

// Bus is a contract to handle commands through a bus
type Bus interface {
	Register(commandType Type, handler Handler)
	// getHandler(commandType Type) Handler
	Dispatch(context.Context, Command) error
}
