package person

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// ErrInvalidPersonID error for invalid person id
var ErrInvalidPersonID = errors.New("invalid Person ID")

// PersonID value object pattern, for person id type
type PersonID struct {
	value string
}

// NewPersonID initializes a new PersonID
func NewPersonID(value string) (PersonID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return PersonID{}, fmt.Errorf("%w: %s", ErrInvalidPersonID, value)
	}

	return PersonID{
		value: v.String(),
	}, nil
}

// String is a toString method.
// Converts the PersonID into string
func (p PersonID) String() string {
	return p.value
}
