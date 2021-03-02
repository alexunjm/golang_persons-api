package personrepository

import (
	"fmt"
	"golang_persons-api/src/domain/person"
)

// PersonRepository handles database actions
type PersonRepository struct {
}

// NewPersonRepository creates a new PersonRepository
func NewPersonRepository() PersonRepository {
	return PersonRepository{}
}

// Save database action
func (PersonRepository) Save(person person.Person) {
	// TODO: save to db
	fmt.Printf("saving person:\n %+v", person)
}

// Update database action
func (PersonRepository) Update(personID person.PersonID, person person.Person) {
	// TODO: update to db
	fmt.Printf("updating person with id <%+v>:\n %+v", personID, person)
}
