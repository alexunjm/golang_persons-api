package query

// Type is a shortcut for string on command types
type Type string

// Query is a contract to retrieve type of contract
type Query interface {
	Type() Type
}
