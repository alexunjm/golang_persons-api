package update

import (
	"context"
	"golang_persons-api/src/domain/module/person"
	"golang_persons-api/src/domain/module/person/domainpersonrepository"
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
	return c.repository.Update(ctx, personModel)
}
