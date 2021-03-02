package personhandler

import "golang_persons-api/src/domain/person"

type CreateCommandHandler struct {
	// createPersonRepository PersonRepository
}

type PersonDTO interface {
	ID() person.PersonID
	Firstname() person.PersonFirstname
	Lastname() person.PersonLastname
	Age() person.PersonAge
}
