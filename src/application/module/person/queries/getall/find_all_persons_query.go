package getall

import "golang_persons-api/src/domain/module/person/query"

var (
	// PersonQueryType is a type for deleting persons query
	PersonQueryType query.Type = "FIND_ALL_PERSONS"
)

// FindAllPersonsQuery struct query all persons
type FindAllPersonsQuery struct{}

// Type returns a query type string
func (FindAllPersonsQuery) Type() query.Type {
	return PersonQueryType
}
