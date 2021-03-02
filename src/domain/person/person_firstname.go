package person

type PersonFirstname struct {
	firstname string
}

func (p PersonFirstname) Firstname() string {
	return p.firstname
}

func NewPersonFirstname(firstname string) PersonFirstname {
	return PersonFirstname{firstname}
}
