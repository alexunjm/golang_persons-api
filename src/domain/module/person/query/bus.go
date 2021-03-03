package query

import (
	"context"
	"golang_persons-api/src/application/module/person/queries"
)

// Bus is a contract to handle commands through a bus
type Bus interface {
	Register(queryType Type, handler Handler)
	// getHandler(commandType Type) Handler
	Dispatch(context.Context, Query) (queries.FindPersonQueryResponse, error)
}
