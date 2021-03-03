package command

import "context"

// Handler is a contract to handle commands
type Handler interface {
	Handle(context.Context, Command) error
}
