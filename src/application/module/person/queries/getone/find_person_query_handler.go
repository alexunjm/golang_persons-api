package getone

import (
	"context"
	"fmt"
	"golang_persons-api/src/application/module/person/queries"
	"golang_persons-api/src/domain/module/person/query"
)

// NewFindPersonQueryHandler initializes a new FindPersonQueryHandler.
func NewFindPersonQueryHandler(service PersonFinderService) FindPersonQueryHandler {
	return FindPersonQueryHandler{service}
}

// FindPersonQueryHandler is a handler for PersonCommand
type FindPersonQueryHandler struct {
	service PersonFinderService
}

// Handle method of query
func (h FindPersonQueryHandler) Handle(ctx context.Context, query query.Query) (queries.FindPersonQueryResponse, error) {
	// casting query to FindPersonQuery
	findPersonQuery, ok := query.(FindPersonQuery)
	if !ok {
		return queries.FindPersonQueryResponse{}, fmt.Errorf("unexpected query FindPersonQuery; found: %+v", query)
	}

	return h.service.Find(
		ctx,
		findPersonQuery.ID,
	)
}
