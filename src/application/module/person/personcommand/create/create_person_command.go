package create

import "golang_persons-api/src/domain/person/command"

var (
	// PersonCommandType is a type for creating persons command
	PersonCommandType command.Type = "CREATE_PERSON"
)

// CreatePersonCommand is a DTO for Person to update
type CreatePersonCommand struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

// Type returns a command type string
func (CreatePersonCommand) Type() command.Type {
	return PersonCommandType
}