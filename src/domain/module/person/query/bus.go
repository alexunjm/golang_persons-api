package query

import (
	"context"
)

// Bus is a contract to handle commands through a bus
type Bus interface {
	Register(queryType Type, handler Handler)
	// getHandler(commandType Type) Handler
	Dispatch(context.Context, Query) (interface{}, error)
}
