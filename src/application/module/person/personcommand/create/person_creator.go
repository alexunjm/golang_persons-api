package create

import (
	"golang_persons-api/src/domain/person"
	"golang_persons-api/src/domain/person/domainpersonrepository"
)

// PersonCreator use case
type PersonCreator struct {
	repository domainpersonrepository.PersonRepository
}

// NewPersonCreator creates a new PersonCreator
func NewPersonCreator(repository domainpersonrepository.PersonRepository) PersonCreator {
	return PersonCreator{
		repository,
	}
}

// Create database action repository call
func (c PersonCreator) Create(id person.PersonID, firstname person.PersonFirstname, lastname person.PersonLastname, age person.PersonAge) {
	personModel := person.NewPersonModel(id, firstname, lastname, age)
	c.repository.Save(personModel)
}
