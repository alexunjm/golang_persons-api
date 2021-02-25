package personcommand

// Updatepersoncommand represents new person to update
type Updatepersoncommand struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

// GetImmutable func returns an immutable command
func (upc Updatepersoncommand) GetImmutable() UpdatepersoncommandContract {
	return immutableUpdatepersoncommand{
		id:        upc.ID,
		firstname: upc.Firstname,
		lastname:  upc.Lastname,
		age:       upc.Age,
	}
}

type immutableUpdatepersoncommand struct {
	id        string
	firstname string
	lastname  string
	age       int
}

func (iupc immutableUpdatepersoncommand) GetID() string {
	return iupc.id
}

func (iupc immutableUpdatepersoncommand) GetFirstname() string {
	return iupc.firstname
}

func (iupc immutableUpdatepersoncommand) GetLastname() string {
	return iupc.lastname
}

func (iupc immutableUpdatepersoncommand) GetAge() int {
	return iupc.age
}
