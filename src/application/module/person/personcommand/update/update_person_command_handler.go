package update

import (
	"fmt"
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
func (h UpdatePersonCommandHandler) Handle(command command.Command) error {
	// casting command to UpdatePersonCommand
	updatePersonCommand, ok := command.(UpdatePersonCommand)
	if !ok {
		return fmt.Errorf("unexpected command UpdatePersonCommand; found: %+v", command)
	}

	return h.service.Update(
		updatePersonCommand.ID,
		updatePersonCommand.Firstname,
		updatePersonCommand.Lastname,
		updatePersonCommand.Age,
	)
}
