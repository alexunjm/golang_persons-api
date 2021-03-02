package person

type PersonAge struct {
	age int
}

func (p PersonAge) Age() int {
	return p.age
}

func NewPersonAge(firstname int) PersonAge {
	return PersonAge{firstname}
}
