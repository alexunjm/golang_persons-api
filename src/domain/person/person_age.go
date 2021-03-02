package person

import (
	"errors"
	"fmt"
	"strconv"
)

// ErrInvalidPersonAge error for invalid person id
var ErrInvalidPersonAge = errors.New("invalid Person Age. Age could be a positive value")

// PersonAge value object pattern, for person age type
type PersonAge struct {
	value string
}

// NewPersonAge initializes a new PersonAge
func NewPersonAge(value int) (PersonAge, error) {
	if value < 0 {
		return PersonAge{}, fmt.Errorf("%w: %d", ErrInvalidPersonAge, value)
	}
	return PersonAge{value: strconv.Itoa(value)}, nil
}

// String is a toString method.
// Converts the PersonAge into string
func (p PersonAge) String() string {
	return p.value
}
