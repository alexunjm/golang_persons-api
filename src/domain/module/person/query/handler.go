package query

import (
	"context"
	"golang_persons-api/src/application/module/person/queries"
)

// Handler is a contract to handle commands
type Handler interface {
	Handle(context.Context, Query) (queries.FindPersonQueryResponse, error)
}
