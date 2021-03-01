package person

type PersonAge struct {
	age string
}

func (p PersonAge) Age() string {
	return p.age
}
