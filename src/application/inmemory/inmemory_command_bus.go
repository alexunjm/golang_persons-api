package inmemory

import (
	"context"
	"fmt"
	"golang_persons-api/src/domain/module/person/command"
)

// CommandBus is a bus for commands, dispatched from request controller
type CommandBus struct {
	handlerMap map[command.Type]command.Handler
}

// NewCommandBus creates a new command bus
func NewCommandBus() command.Bus {
	return CommandBus{
		handlerMap: make(map[command.Type]command.Handler),
	}
}

// Register adds handler for type commandType
func (b CommandBus) Register(commandType command.Type, handler command.Handler) {
	b.handlerMap[commandType] = handler
}

func (b CommandBus) getHandler(commandType command.Type) command.Handler {
	handler, ok := b.handlerMap[commandType]
	if !ok {
		panic(fmt.Sprintf("unregistered handler for %v command type", commandType))
	}
	return handler
}

// Dispatch finds handler for command and delegates to handler, handle the command
func (b CommandBus) Dispatch(ctx context.Context, command command.Command) error {
	handler := b.getHandler(command.Type())
	return handler.Handle(ctx, command)
}
