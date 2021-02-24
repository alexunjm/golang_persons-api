package person_query

// FindPersonQuery represents a person to find
type FindPersonQuery struct {
	ID string `json:"id"`
}

// GetImmutable func returns an immutable query
func (fpq FindPersonQuery) GetImmutable() FindPersonQueryContract {
	return immutableFindPersonQuery{
		id: fpq.ID,
	}
}

type immutableFindPersonQuery struct {
	id string
}

func (ifpq immutableFindPersonQuery) GetID() string {
	return ifpq.id
}
