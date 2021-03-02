package update

import (
	"golang_persons-api/src/domain/person"
	"golang_persons-api/src/domain/person/domainpersonrepository"
)

// PersonUpdater use case
type PersonUpdater struct {
	repository domainpersonrepository.PersonRepository
}

// NewPersonUpdater creates a new PersonUpdater
func NewPersonUpdater(repository domainpersonrepository.PersonRepository) PersonUpdater {
	return PersonUpdater{repository}
}

// Update database action repository call
func (c PersonUpdater) Update(srcPersonID person.PersonID, id person.PersonID, firstname person.PersonFirstname, lastname person.PersonLastname, age person.PersonAge) {
	personModel := person.NewPersonModel(id, firstname, lastname, age)
	c.repository.Update(srcPersonID, personModel)
}
