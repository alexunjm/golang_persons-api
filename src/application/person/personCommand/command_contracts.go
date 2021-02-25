package personCommand

// CreatePersonCommandContract defines command getter methods
type CreatePersonCommandContract interface {
	GetID() string
	GetFirstname() string
	GetLastname() string
	GetAge() int
}

// UpdatePersonCommandContract defines command getter methods
type UpdatePersonCommandContract interface {
	GetID() string
	GetFirstname() string
	GetLastname() string
	GetAge() int
}

// DeletePersonCommandContract defines command getter methods
type DeletePersonCommandContract interface {
	GetID() string
}
