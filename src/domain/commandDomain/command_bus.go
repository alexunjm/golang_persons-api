package commandDomain

type CommandBus interface {
	dispatch() string
}
