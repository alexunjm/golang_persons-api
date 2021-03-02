package update

import "golang_persons-api/src/domain/person/command"

var (
	// PersonCommandType is a type for updating persons command
	PersonCommandType command.Type = "UPDATE_PERSON"
)

// UpdatePersonCommand is a DTO for Person to update
type UpdatePersonCommand struct {
	srcID     string `json:"src_id"`
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

// Type returns a command type string
func (UpdatePersonCommand) Type() command.Type {
	return PersonCommandType
}
