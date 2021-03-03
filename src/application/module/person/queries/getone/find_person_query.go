package getone

import "golang_persons-api/src/domain/module/person/query"

var (
	// PersonQueryType is a type for deleting persons query
	PersonQueryType query.Type = "FIND_PERSON"
)

// FindPersonQuery is a handler for query a person
type FindPersonQuery struct {
	ID string
}

// Type returns a query type string
func (FindPersonQuery) Type() query.Type {
	return PersonQueryType
}
