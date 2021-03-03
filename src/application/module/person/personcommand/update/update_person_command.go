package update

import "golang_persons-api/src/domain/module/person/command"

var (
	// PersonCommandType is a type for updating persons command
	PersonCommandType command.Type = "UPDATE_PERSON"
)

// UpdatePersonCommand is a DTO for Person to update
type UpdatePersonCommand struct {
	ID        string
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Age       int    `json:"age" binding:"required"`
}

// Type returns a command type string
func (UpdatePersonCommand) Type() command.Type {
	return PersonCommandType
}
