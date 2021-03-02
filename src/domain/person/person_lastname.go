package person

type PersonLastname struct {
	lastname string
}

func (p PersonLastname) Lastname() string {
	return p.lastname
}

func NewPersonLastname(firstname string) PersonLastname {
	return PersonLastname{firstname}
}
