package update

import (
	"fmt"
	"golang_persons-api/src/domain/person"
	"golang_persons-api/src/domain/person/command"
)

// NewUpdatePersonCommandHandler initializes a new UpdatePersonCommandHandler.
func NewUpdatePersonCommandHandler(service PersonUpdaterService) UpdatePersonCommandHandler {
	return UpdatePersonCommandHandler{service}
}

// UpdatePersonCommandHandler is a handler for PersonCommand
type UpdatePersonCommandHandler struct {
	service PersonUpdaterService
}

// Handle method of command
func (h UpdatePersonCommandHandler) Handle(command command.Command) {
	// casting command to UpdatePersonCommand
	updatePersonCommand, ok := command.(UpdatePersonCommand)
	if !ok {
		panic(fmt.Sprintf("unexpected command updatePersonCommand; found: %+v", command))
	}

	srcID := person.NewPersonID(updatePersonCommand.srcID)
	id := person.NewPersonID(updatePersonCommand.ID)
	firstname := person.NewPersonFirstname(updatePersonCommand.Firstname)
	lastname := person.NewPersonLastname(updatePersonCommand.Lastname)
	age := person.NewPersonAge(updatePersonCommand.Age)

	h.service.Update(srcID, id, firstname, lastname, age)
}
