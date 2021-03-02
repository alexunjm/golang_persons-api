package create

import "golang_persons-api/src/domain/person"

// PersonCreator use case
type PersonCreator struct {
}

// NewPersonCreator creates a new PersonCreator
func NewPersonCreator() PersonCreator {
	return PersonCreator{}
}

func (c PersonCreator) Create(id person.PersonID, firstname person.PersonFirstname, lastname person.PersonLastname, age person.PersonAge) {
	// TODO: Create a person domain
}
