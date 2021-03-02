package person

// Person domain model
type Person struct {
	id        PersonID
	firstname PersonFirstname
	lastname  PersonLastname
	age       PersonAge
}

// NewPersonModel creates a new person
func NewPersonModel(id PersonID, firstname PersonFirstname, lastname PersonLastname, age PersonAge) Person {
	return Person{
		id,
		firstname,
		lastname,
		age,
	}
}
