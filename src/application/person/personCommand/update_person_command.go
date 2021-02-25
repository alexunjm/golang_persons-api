package personCommand

// UpdatePersonCommand represents new person to update
type UpdatePersonCommand struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

// GetImmutable func returns an immutable command
func (upc UpdatePersonCommand) GetImmutable() UpdatePersonCommandContract {
	return immutableUpdatePersonCommand{
		id:        upc.ID,
		firstname: upc.Firstname,
		lastname:  upc.Lastname,
		age:       upc.Age,
	}
}

type immutableUpdatePersonCommand struct {
	id        string
	firstname string
	lastname  string
	age       int
}

func (iupc immutableUpdatePersonCommand) GetID() string {
	return iupc.id
}

func (iupc immutableUpdatePersonCommand) GetFirstname() string {
	return iupc.firstname
}

func (iupc immutableUpdatePersonCommand) GetLastname() string {
	return iupc.lastname
}

func (iupc immutableUpdatePersonCommand) GetAge() int {
	return iupc.age
}
