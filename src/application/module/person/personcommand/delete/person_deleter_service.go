package delete

import (
	"context"
	"golang_persons-api/src/domain/person"
	"golang_persons-api/src/domain/person/domainpersonrepository"
)

// PersonDeleterService (use case)
type PersonDeleterService struct {
	repository domainpersonrepository.PersonRepository
}

// NewPersonDeleter creates a new PersonDeleterService
func NewPersonDeleter(repository domainpersonrepository.PersonRepository) PersonDeleterService {
	return PersonDeleterService{repository}
}

// Delete database action repository call
func (c PersonDeleterService) Delete(ctx context.Context, id string) error {
	personID, err := person.NewPersonID(id)
	if err != nil {
		return err
	}
	return c.repository.Delete(ctx, personID)
}
