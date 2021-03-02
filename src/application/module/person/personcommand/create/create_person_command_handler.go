package create

import (
	"fmt"
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
func (h CreatePersonCommandHandler) Handle(command command.Command) error {
	// casting command to CreatePersonCommand
	createPersonCommand, ok := command.(CreatePersonCommand)
	if !ok {
		return fmt.Errorf("unexpected command CreatePersonCommand; found: %+v", command)
	}

	return h.service.Create(
		createPersonCommand.ID,
		createPersonCommand.Firstname,
		createPersonCommand.Lastname,
		createPersonCommand.Age,
	)
}
