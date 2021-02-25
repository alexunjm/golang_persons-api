package personcommand

// CreatepersoncommandContract defines command getter methods
type CreatepersoncommandContract interface {
	GetID() string
	GetFirstname() string
	GetLastname() string
	GetAge() int
}

// UpdatepersoncommandContract defines command getter methods
type UpdatepersoncommandContract interface {
	GetID() string
	GetFirstname() string
	GetLastname() string
	GetAge() int
}

// DeletepersoncommandContract defines command getter methods
type DeletepersoncommandContract interface {
	GetID() string
}
