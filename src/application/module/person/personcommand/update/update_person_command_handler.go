package update

import (
	"golang_persons-api/src/domain/person"
	"golang_persons-api/src/domain/person/command"
)

// NewUpdatePersonCommandHandler initializes a new UpdatePersonCommandHandler.
func NewUpdatePersonCommandHandler(creator PersonUpdater) UpdatePersonCommandHandler {
	return UpdatePersonCommandHandler{
		creator: creator,
	}
}

// UpdatePersonCommandHandler is a handler for PersonCommand
type UpdatePersonCommandHandler struct {
	creator PersonUpdater
}

// Handle method of command
func (h UpdatePersonCommandHandler) Handle(command command.Command) {
	// casting command to UpdatePersonCommand
	personCommand := command.(UpdatePersonCommand)

	id := person.NewPersonID(personCommand.ID)
	firstname := person.NewPersonFirstname(personCommand.Firstname)
	lastname := person.NewPersonLastname(personCommand.Lastname)
	age := person.NewPersonAge(personCommand.Age)

	h.creator.Update(id, firstname, lastname, age)
}
