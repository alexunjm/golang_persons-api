package create

import (
	"context"
	"golang_persons-api/src/domain/module/person"
	"golang_persons-api/src/domain/module/person/domainpersonrepository"
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
func (c PersonCreatorService) Create(ctx context.Context, id string, firstname string, lastname string, age int) error {

	personModel, err := person.NewPersonModel(id, firstname, lastname, age)
	if err != nil {
		return err
	}

	err = c.repository.Save(ctx, personModel)
	if err != nil {
		return err
	}

	return nil
}
