package personcommand

// Createpersoncommand represents new person to create
type Createpersoncommand struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

// GetImmutable func returns an immutable command
func (cpc Createpersoncommand) GetImmutable() CreatepersoncommandContract {
	return immutableCreatepersoncommand{
		id:        cpc.ID,
		firstname: cpc.Firstname,
		lastname:  cpc.Lastname,
		age:       cpc.Age,
	}
}

type immutableCreatepersoncommand struct {
	id        string
	firstname string
	lastname  string
	age       int
}

func (icpc immutableCreatepersoncommand) GetID() string {
	return icpc.id
}

func (icpc immutableCreatepersoncommand) GetFirstname() string {
	return icpc.firstname
}

func (icpc immutableCreatepersoncommand) GetLastname() string {
	return icpc.lastname
}

func (icpc immutableCreatepersoncommand) GetAge() int {
	return icpc.age
}
