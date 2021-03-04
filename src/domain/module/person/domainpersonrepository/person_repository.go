package domainpersonrepository

import (
	"context"
	"golang_persons-api/src/domain/module/person"
)

// PersonRepository is responsible for persisting person data model
type PersonRepository interface {
	Save(ctx context.Context, person person.Person) error
	Update(ctx context.Context, person person.Person) error
	Delete(ctx context.Context, personID person.PersonID) error
	Find(ctx context.Context, personID person.PersonID) (interface{}, error)
	// FindAll(ctx context.Context) (person.Persons, error)
}
