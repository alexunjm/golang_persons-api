package person

type PersonLastname struct {
	lastname string
}

func (p PersonLastname) Lastname() string {
	return p.lastname
}
