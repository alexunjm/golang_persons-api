package person

type PersonId struct {
	id string
}

func (p PersonId) Id() string {
	return p.id
}
