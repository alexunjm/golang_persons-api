package api

/*
type BusContract interface {
	dispatch(command CommandContract)
}

type CommandContract interface{}
type CommandBusContract interface {
	dispatch(command CommandContract)
}

type QueryBusContract interface{}

type ApiHttpStatusCodeExceptions interface{}

type ApiControllerContract interface {
	// GetQueryBus() QueryBusContract
	// GetCommandBus() CommandBusContract
	// GetExceptionHandler() ApiHttpStatusCodeExceptions
	Dispatch(command CommandContract)
}

type ApiController struct {
	QueryBus         QueryBusContract
	CommandBus       CommandBusContract
	ExceptionHandler ApiHttpStatusCodeExceptions
}

func (ac ApiController) Dispatch(command CommandContract) {
	ac.CommandBus.dispatch(command)
}

type SyncCommandBus struct {
	Bus BusContract
}

func (syncCB SyncCommandBus) dispatch(command CommandContract) {
	syncCB.Bus.dispatch(command)
}
*/
