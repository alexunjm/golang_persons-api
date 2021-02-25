package personCommand

// CreatePersonCommand represents new person to create
type CreatePersonCommand struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

// GetImmutable func returns an immutable command
func (cpc CreatePersonCommand) GetImmutable() CreatePersonCommandContract {
	return immutableCreatePersonCommand{
		id:        cpc.ID,
		firstname: cpc.Firstname,
		lastname:  cpc.Lastname,
		age:       cpc.Age,
	}
}

type immutableCreatePersonCommand struct {
	id        string
	firstname string
	lastname  string
	age       int
}

func (icpc immutableCreatePersonCommand) GetID() string {
	return icpc.id
}

func (icpc immutableCreatePersonCommand) GetFirstname() string {
	return icpc.firstname
}

func (icpc immutableCreatePersonCommand) GetLastname() string {
	return icpc.lastname
}

func (icpc immutableCreatePersonCommand) GetAge() int {
	return icpc.age
}
