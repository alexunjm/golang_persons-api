package personquery

// Findpersonquery represents a person to find
type Findpersonquery struct {
	ID string `json:"id"`
}

// GetImmutable func returns an immutable query
func (fpq Findpersonquery) GetImmutable() FindpersonqueryContract {
	return immutableFindpersonquery{
		id: fpq.ID,
	}
}

type immutableFindpersonquery struct {
	id string
}

func (ifpq immutableFindpersonquery) GetID() string {
	return ifpq.id
}
