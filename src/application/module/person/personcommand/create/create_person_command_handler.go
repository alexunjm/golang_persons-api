package create

import (
	"fmt"
	"golang_persons-api/src/domain/person"
	"golang_persons-api/src/domain/person/command"
)

// NewCreatePersonCommandHandler initializes a new CreatePersonCommandHandler.
func NewCreatePersonCommandHandler(service PersonCreatorService) CreatePersonCommandHandler {
	return CreatePersonCommandHandler{service}
}

// CreatePersonCommandHandler is a handler for PersonCommand
type CreatePersonCommandHandler struct {
	service PersonCreatorService
}

// Handle method of command
func (h CreatePersonCommandHandler) Handle(command command.Command) {
	// casting command to CreatePersonCommand
	createPersonCommand, ok := command.(CreatePersonCommand)
	if !ok {
		panic(fmt.Sprintf("unexpected command CreatePersonCommand; found: %+v", command))
	}

	id := person.NewPersonID(createPersonCommand.ID)
	firstname := person.NewPersonFirstname(createPersonCommand.Firstname)
	lastname := person.NewPersonLastname(createPersonCommand.Lastname)
	age := person.NewPersonAge(createPersonCommand.Age)

	h.service.Create(id, firstname, lastname, age)
}
