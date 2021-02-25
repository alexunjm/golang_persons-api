package personcommand

// Deletepersoncommand represents a person to delete
type Deletepersoncommand struct {
	ID string `json:"id"`
}

// GetImmutable func returns an immutable command
func (dpc Deletepersoncommand) GetImmutable() DeletepersoncommandContract {
	return immutableDeletepersoncommand{
		id: dpc.ID,
	}
}

type immutableDeletepersoncommand struct {
	id string
}

func (idpc immutableDeletepersoncommand) GetID() string {
	return idpc.id
}
