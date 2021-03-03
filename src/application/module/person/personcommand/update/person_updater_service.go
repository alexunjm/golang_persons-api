package update

import (
	"context"
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
func (c PersonUpdaterService) Update(ctx context.Context, id string, firstname string, lastname string, age int) error {
	personModel, err := person.NewPersonModel(id, firstname, lastname, age)
	if err != nil {
		return err
	}
	c.repository.Update(ctx, personModel)
	return nil
}
