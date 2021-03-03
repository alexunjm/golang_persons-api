package person

// Person domain model
type Person struct {
	id        PersonID
	firstname PersonFirstname
	lastname  PersonLastname
	age       PersonAge
}

// NewPersonModel creates a new person
func NewPersonModel(pID string, pFirstname string, pLastname string, pAge int) (Person, error) {

	var (
		id        PersonID
		firstname PersonFirstname
		lastname  PersonLastname
		age       PersonAge
		err       error
	)

	if id, err = NewPersonID(pID); err != nil {
		return Person{}, err
	}

	if firstname, err = NewPersonFirstname(pFirstname); err != nil {
		return Person{}, err
	}

	if lastname, err = NewPersonLastname(pLastname); err != nil {
		return Person{}, err
	}

	if age, err = NewPersonAge(pAge); err != nil {
		return Person{}, err
	}

	return Person{
		id,
		firstname,
		lastname,
		age,
	}, nil
}

// ID returns PersonID
func (p Person) ID() PersonID {
	return p.id
}

// Firstname returns PersonFirstname
func (p Person) Firstname() PersonFirstname {
	return p.firstname
}

// Lastname returns PersonLastname
func (p Person) Lastname() PersonLastname {
	return p.lastname
}

// Age returns PersonAge
func (p Person) Age() PersonAge {
	return p.age
}
