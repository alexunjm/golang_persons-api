package commanddomain

type CommandBus interface {
	dispatch() string
}
