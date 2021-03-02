package update

import (
	"golang_persons-api/src/domain/person"
	"golang_persons-api/src/domain/person/command"
)

// NewUpdatePersonCommandHandler initializes a new UpdatePersonCommandHandler.
func NewUpdatePersonCommandHandler(updater PersonUpdater) UpdatePersonCommandHandler {
	return UpdatePersonCommandHandler{updater}
}

// UpdatePersonCommandHandler is a handler for PersonCommand
type UpdatePersonCommandHandler struct {
	updater PersonUpdater
}

// Handle method of command
func (h UpdatePersonCommandHandler) Handle(command command.Command) {
	// casting command to UpdatePersonCommand
	updatePersonCommand := command.(UpdatePersonCommand)

	srcID := person.NewPersonID(updatePersonCommand.srcID)
	id := person.NewPersonID(updatePersonCommand.ID)
	firstname := person.NewPersonFirstname(updatePersonCommand.Firstname)
	lastname := person.NewPersonLastname(updatePersonCommand.Lastname)
	age := person.NewPersonAge(updatePersonCommand.Age)

	h.updater.Update(srcID, id, firstname, lastname, age)
}
