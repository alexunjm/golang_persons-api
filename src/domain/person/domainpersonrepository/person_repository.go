package domainpersonrepository

import (
	"context"
	"golang_persons-api/src/domain/person"
)

// PersonRepository is responsible for persisting person data model
type PersonRepository interface {
	Save(ctx context.Context, person person.Person) error
	Update(ctx context.Context, person person.Person) error
}
