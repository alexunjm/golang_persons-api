package domainpersonrepository

import "golang_persons-api/src/domain/person"

// PersonRepository is responsible for persisting person data model
type PersonRepository interface {
	Save(person person.Person)
	Update(person person.Person)
}
