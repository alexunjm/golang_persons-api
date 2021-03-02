package create

import (
	"golang_persons-api/src/domain/person"
	"golang_persons-api/src/domain/person/domainpersonrepository"
)

// PersonCreatorService (use case)
type PersonCreatorService struct {
	repository domainpersonrepository.PersonRepository
}

// NewPersonCreator creates a new PersonCreatorService
func NewPersonCreator(repository domainpersonrepository.PersonRepository) PersonCreatorService {
	return PersonCreatorService{
		repository,
	}
}

// Create database action repository call
func (c PersonCreatorService) Create(id string, firstname string, lastname string, age int) error {

	personModel, err := person.NewPersonModel(id, firstname, lastname, age)
	if err != nil {
		return err
	}
	c.repository.Save(personModel)
	return nil
}
