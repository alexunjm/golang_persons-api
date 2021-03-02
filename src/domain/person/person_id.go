package person

type PersonID struct {
	id string
}

func (p PersonID) Id() string {
	return p.id
}

func NewPersonID(id string) PersonID {
	return PersonID{id}
}
