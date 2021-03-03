package delete

import (
	"context"
	"fmt"
	"golang_persons-api/src/domain/module/person/command"
)

// NewDeletePersonCommandHandler initializes a new DeletePersonCommandHandler.
func NewDeletePersonCommandHandler(service PersonDeleterService) DeletePersonCommandHandler {
	return DeletePersonCommandHandler{service}
}

// DeletePersonCommandHandler is a handler for PersonCommand
type DeletePersonCommandHandler struct {
	service PersonDeleterService
}

// Handle method of command
func (h DeletePersonCommandHandler) Handle(ctx context.Context, command command.Command) error {
	// casting command to DeletePersonCommand
	deletePersonCommand, ok := command.(DeletePersonCommand)
	if !ok {
		return fmt.Errorf("unexpected command DeletePersonCommand; found: %+v", command)
	}

	return h.service.Delete(
		ctx,
		deletePersonCommand.ID,
	)
}
