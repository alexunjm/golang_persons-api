package person_command

// DeletePersonCommand represents a person to delete
type DeletePersonCommand struct {
	ID string `json:"id"`
}

// GetImmutable func returns an immutable command
func (dpc DeletePersonCommand) GetImmutable() DeletePersonCommandContract {
	return immutableDeletePersonCommand{
		id: dpc.ID,
	}
}

type immutableDeletePersonCommand struct {
	id string
}

func (idpc immutableDeletePersonCommand) GetID() string {
	return idpc.id
}
