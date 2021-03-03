package command

// Type is a shortcut for string on command types
type Type string

// Command is a contract to retrieve type of contract
type Command interface {
	Type() Type
}
