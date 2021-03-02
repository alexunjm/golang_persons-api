package create

import "golang_persons-api/src/domain/person/command"

var (
	// PersonCommandType is a type for creating persons command
	PersonCommandType command.Type = "CREATE_PERSON"
)

type CreatePersonCommand struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func (CreatePersonCommand) Type() command.Type {
	return PersonCommandType
}
