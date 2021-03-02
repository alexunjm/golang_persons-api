package create

import (
	"golang_persons-api/src/domain/person"
	"golang_persons-api/src/domain/person/command"
)

// NewCreatePersonCommandHandler initializes a new CreatePersonCommandHandler.
func NewCreatePersonCommandHandler(creator PersonCreator) CreatePersonCommandHandler {
	return CreatePersonCommandHandler{
		creator: creator,
	}
}

// CreatePersonCommandHandler is a handler for PersonCommand
type CreatePersonCommandHandler struct {
	creator PersonCreator
}

// Handle method of command
func (h CreatePersonCommandHandler) Handle(command command.Command) {
	// casting command to CreatePersonCommand
	createPersonCommand := command.(CreatePersonCommand)

	id := person.NewPersonID(createPersonCommand.ID)
	firstname := person.NewPersonFirstname(createPersonCommand.Firstname)
	lastname := person.NewPersonLastname(createPersonCommand.Lastname)
	age := person.NewPersonAge(createPersonCommand.Age)

	h.creator.Create(id, firstname, lastname, age)
}
