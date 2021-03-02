package domainpersonrepository

import "golang_persons-api/src/domain/person"

// PersonRepository is responsible for persisting person data model
type PersonRepository interface {
	save(person person.Person)
	update(srcPersonID string, person person.Person)
}
