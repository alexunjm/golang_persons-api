package getall

import (
	"context"
	"fmt"
	"golang_persons-api/src/application/module/person/queries"
	"golang_persons-api/src/domain/module/person/query"
)

// NewFindAllPersonsQueryHandler initializes a new FindAllPersonsQueryHandler.
func NewFindAllPersonsQueryHandler(service AllPersonsFinderService) FindAllPersonsQueryHandler {
	return FindAllPersonsQueryHandler{service}
}

// FindAllPersonsQueryHandler is a handler for PersonCommand
type FindAllPersonsQueryHandler struct {
	service AllPersonsFinderService
}

// Handle method of query
func (h FindAllPersonsQueryHandler) Handle(ctx context.Context, query query.Query) (interface{}, error) {
	// casting query to FindAllPersonsQuery
	findAllPersonsQuery, ok := query.(FindAllPersonsQuery)
	if !ok {
		return queries.FindAllPersonsQueryResponse{}, fmt.Errorf("unexpected query FindAllPersonsQuery; found: %+v", query)
	}
	doNothing(findAllPersonsQuery)

	return h.service.FindAll(
		ctx,
	)
}

func doNothing(_ FindAllPersonsQuery) {}
