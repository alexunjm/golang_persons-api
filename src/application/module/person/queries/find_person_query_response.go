package queries

// FindPersonQueryResponse is a handler for query a person
type FindPersonQueryResponse struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

// FindAllPersonsQueryResponse is a handler for query a person
type FindAllPersonsQueryResponse []FindPersonQueryResponse
