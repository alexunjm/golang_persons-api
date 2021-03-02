package persondependencyinjection

import (
	"golang_persons-api/src/application/module/person/personcommand/create"
	"golang_persons-api/src/application/module/person/personcommand/update"
	"golang_persons-api/src/domain/person/command"
)

// RegisterPersonCommandHandlers func registers all command handlers
func RegisterPersonCommandHandlers(bus command.Bus) {

	personCreator := create.NewPersonCreator()
	createPersonCommandHandler := create.NewCreatePersonCommandHandler(personCreator)
	bus.Register(create.PersonCommandType, createPersonCommandHandler)

	personUpdater := update.NewPersonUpdater()
	updatePersonCommandHandler := update.NewUpdatePersonCommandHandler(personUpdater)
	bus.Register(update.PersonCommandType, updatePersonCommandHandler)
}
