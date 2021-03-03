package query

import (
	"context"
	"golang_persons-api/src/domain/module/person"
)

// Bus is a contract to handle commands through a bus
type Bus interface {
	Register(commandType Type, handler Handler)
	// getHandler(commandType Type) Handler
	Dispatch(context.Context, Query) (person.Person, error)
}
