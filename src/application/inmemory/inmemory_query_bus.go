package inmemory

import (
	"context"
	"fmt"
	"golang_persons-api/src/domain/module/person/query"
)

// QueryBus is a bus for queries, dispatched from request controller
type QueryBus struct {
	handlerMap map[query.Type]query.Handler
}

// NewQueryBus creates a new query bus
func NewQueryBus() query.Bus {
	return QueryBus{
		handlerMap: make(map[query.Type]query.Handler),
	}
}

// Register adds handler for type queryType
func (b QueryBus) Register(queryType query.Type, handler query.Handler) {
	b.handlerMap[queryType] = handler
}

func (b QueryBus) getHandler(queryType query.Type) query.Handler {
	handler, ok := b.handlerMap[queryType]
	if !ok {
		panic(fmt.Sprintf("unregistered handler for %v query type", queryType))
	}
	return handler
}

// Dispatch finds handler for query and delegates to handler, handle the query
func (b QueryBus) Dispatch(ctx context.Context, query query.Query) (interface{}, error) {
	handler := b.getHandler(query.Type())
	return handler.Handle(ctx, query)
}
