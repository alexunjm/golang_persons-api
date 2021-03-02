package person

import (
	"errors"
	"strings"
)

// ErrEmptyPersonFirstname error for invalid person id
var ErrEmptyPersonFirstname = errors.New("the field person firstname can not be empty")

// PersonFirstname value object pattern, for person firstname type
type PersonFirstname struct {
	value string
}

// NewPersonFirstname initializes a new PersonFirstname
func NewPersonFirstname(value string) (PersonFirstname, error) {
	value = strings.Trim(value, " ")
	if value == "" {
		return PersonFirstname{}, ErrEmptyPersonFirstname
	}
	return PersonFirstname{value}, nil
}

// String is a toString method.
// Converts the PersonFirstname into string
func (p PersonFirstname) String() string {
	return p.value
}
