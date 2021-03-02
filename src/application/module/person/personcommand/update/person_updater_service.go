package update

import (
	"golang_persons-api/src/domain/person"
	"golang_persons-api/src/domain/person/domainpersonrepository"
)

// PersonUpdaterService (use case)
type PersonUpdaterService struct {
	repository domainpersonrepository.PersonRepository
}

// NewPersonUpdater creates a new PersonUpdaterService
func NewPersonUpdater(repository domainpersonrepository.PersonRepository) PersonUpdaterService {
	return PersonUpdaterService{repository}
}

// Update database action repository call
func (c PersonUpdaterService) Update(srcPersonID person.PersonID, id person.PersonID, firstname person.PersonFirstname, lastname person.PersonLastname, age person.PersonAge) {
	personModel := person.NewPersonModel(id, firstname, lastname, age)
	c.repository.Update(srcPersonID, personModel)
}
