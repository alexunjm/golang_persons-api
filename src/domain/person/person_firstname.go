package person

type PersonFirstname struct {
	firstname string
}

func (p PersonFirstname) Firstname() string {
	return p.firstname
}
