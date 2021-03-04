package getall

import (
	"context"
	"golang_persons-api/src/domain/module/person/domainpersonrepository"
)

// AllPersonsFinderService (use case)
type AllPersonsFinderService struct {
	repository domainpersonrepository.PersonRepository
}

// NewAllPersonsFinder creates a new AllPersonsFinderService
func NewAllPersonsFinder(repository domainpersonrepository.PersonRepository) AllPersonsFinderService {
	return AllPersonsFinderService{
		repository,
	}
}

// FindAll database action repository call
func (c AllPersonsFinderService) FindAll(ctx context.Context) ([]interface{}, error) {

	return c.repository.FindAll(ctx)
}
