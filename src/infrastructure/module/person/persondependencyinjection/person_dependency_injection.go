package persondependencyinjection

import (
	"database/sql"
	"golang_persons-api/src/application/module/person/personcommand/create"
	"golang_persons-api/src/application/module/person/personcommand/update"
	"golang_persons-api/src/domain/person/command"
	"golang_persons-api/src/infrastructure/module/person/personrepository"
	"time"
)

// RegisterPersonCommandHandlers func registers all command handlers
func RegisterPersonCommandHandlers(bus command.Bus, db *sql.DB, dbTimeout time.Duration) {

	personRepository := personrepository.NewPersonRepository(db, dbTimeout)

	personCreator := create.NewPersonCreator(personRepository)
	createPersonCommandHandler := create.NewCreatePersonCommandHandler(personCreator)
	bus.Register(create.PersonCommandType, createPersonCommandHandler)

	personUpdater := update.NewPersonUpdater(personRepository)
	updatePersonCommandHandler := update.NewUpdatePersonCommandHandler(personUpdater)
	bus.Register(update.PersonCommandType, updatePersonCommandHandler)
}
