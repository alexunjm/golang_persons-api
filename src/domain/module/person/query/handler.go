package query

import (
	"context"
)

// Handler is a contract to handle commands
type Handler interface {
	Handle(context.Context, Query) (interface{}, error)
}
