package delete

import "golang_persons-api/src/domain/person/command"

var (
	// PersonCommandType is a type for deleting persons command
	PersonCommandType command.Type = "DELETE_PERSON"
)

// DeletePersonCommand is a DTO for Person to delete
type DeletePersonCommand struct {
	ID string
}

// Type returns a command type string
func (DeletePersonCommand) Type() command.Type {
	return PersonCommandType
}
