package command_domain

type CommandBus interface {
	dispatch() string
}
