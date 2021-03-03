package personrepository

import (
	"context"
	"fmt"
	"golang_persons-api/src/domain/person"

	"github.com/huandu/go-sqlbuilder"
)

// PersonRepository handles database actions
type PersonRepository struct {
}

// NewPersonRepository creates a new PersonRepository
func NewPersonRepository() PersonRepository {
	return PersonRepository{}
}

// Save database action
func (PersonRepository) Save(person person.Person) error {
	// TODO: save to db
	fmt.Printf("saving person:\n %+v", person)

	personSQLStruct := sqlbuilder.NewStruct(new(sqlPerson))
	query, args := personSQLStruct.InsertInto(sqlPersonTable, sqlPerson{
		ID:        person.ID().String(),
		Firstname: person.Firstname().String(),
		Lastname:  person.Lastname().String(),
		Age:       person.Age().Int(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(context.Background(), r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist person on database: %v", err)
	}

	return nil
}

// Update database action
func (PersonRepository) Update(person person.Person) error {
	// TODO: update to db
	fmt.Printf("updating person :\n %+v", person)
	return nil
}
