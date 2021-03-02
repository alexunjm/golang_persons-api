package update

import "golang_persons-api/src/domain/person"

// PersonUpdater use case
type PersonUpdater struct {
}

// NewPersonUpdater creates a new PersonUpdater
func NewPersonUpdater() PersonUpdater {
	return PersonUpdater{}
}

func (c PersonUpdater) Update(id person.PersonID, firstname person.PersonFirstname, lastname person.PersonLastname, age person.PersonAge) {
	// TODO: Update a person domain
}
