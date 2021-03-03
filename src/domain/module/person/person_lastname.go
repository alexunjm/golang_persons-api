package person

import (
	"errors"
	"strings"
)

// ErrEmptyPersonLastname error for invalid person id
var ErrEmptyPersonLastname = errors.New("the field person lastname can not be empty")

// PersonLastname value object pattern, for person lastname type
type PersonLastname struct {
	value string
}

// NewPersonLastname initializes a new PersonLastname
func NewPersonLastname(value string) (PersonLastname, error) {
	value = strings.Trim(value, " ")
	if value == "" {
		return PersonLastname{}, ErrEmptyPersonLastname
	}
	return PersonLastname{value}, nil
}

// String is a toString method.
// Converts the PersonLastname into string
func (p PersonLastname) String() string {
	return p.value
}
