package query

import (
	"context"
	"golang_persons-api/src/domain/module/person"
)

// Handler is a contract to handle commands
type Handler interface {
	Handle(context.Context, Query) (person.Person, error)
}
