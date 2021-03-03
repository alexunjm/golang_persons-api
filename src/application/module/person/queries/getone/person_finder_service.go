package getone

import (
	"context"
	"golang_persons-api/src/domain/module/person"
	"golang_persons-api/src/domain/module/person/domainpersonrepository"
)

// PersonFinderService (use case)
type PersonFinderService struct {
	repository domainpersonrepository.PersonRepository
}

// NewPersonFinder creates a new PersonFinderService
func NewPersonFinder(repository domainpersonrepository.PersonRepository) PersonFinderService {
	return PersonFinderService{
		repository,
	}
}

// Find database action repository call
func (c PersonFinderService) Find(ctx context.Context, id string) (person.Person, error) {

	personID, err := person.NewPersonID(id)
	if err != nil {
		return person.Person{}, err
	}

	return c.repository.Find(ctx, personID)
}
