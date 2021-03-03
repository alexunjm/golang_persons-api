package persondependencyinjection

import (
	"database/sql"
	"golang_persons-api/src/application/module/person/commands/create"
	"golang_persons-api/src/application/module/person/commands/delete"
	"golang_persons-api/src/application/module/person/commands/update"
	"golang_persons-api/src/application/module/person/queries/getone"
	"golang_persons-api/src/domain/module/person/command"
	"golang_persons-api/src/domain/module/person/query"
	repository "golang_persons-api/src/infrastructure/module/person/personrepository"
	"time"
)

// RegisterPersonCommandHandlers func registers all command handlers
func RegisterPersonCommandHandlers(bus command.Bus, db *sql.DB, dbTimeout time.Duration) {

	personRepository := repository.NewPersonRepository(db, dbTimeout)

	personCreator := create.NewPersonCreator(personRepository)
	createPersonCommandHandler := create.NewCreatePersonCommandHandler(personCreator)
	bus.Register(create.PersonCommandType, createPersonCommandHandler)

	personUpdater := update.NewPersonUpdater(personRepository)
	updatePersonCommandHandler := update.NewUpdatePersonCommandHandler(personUpdater)
	bus.Register(update.PersonCommandType, updatePersonCommandHandler)

	personDeleter := delete.NewPersonDeleter(personRepository)
	deletePersonCommandHandler := delete.NewDeletePersonCommandHandler(personDeleter)
	bus.Register(delete.PersonCommandType, deletePersonCommandHandler)
}

// RegisterPersonQueryHandlers func registers all cqueryd handlers
func RegisterPersonQueryHandlers(bus query.Bus, db *sql.DB, dbTimeout time.Duration) {

	personRepository := repository.NewPersonRepository(db, dbTimeout)

	personFinder := getone.NewPersonFinder(personRepository)
	findPersonQueryHandler := getone.NewFindPersonQueryHandler(personFinder)
	bus.Register(getone.PersonQueryType, findPersonQueryHandler)
	/*
		// TODO: find all
		findAllPersons := getall.NewFindAllPersons(personRepository)
		findAllPersonsQueryHandler := getall.NewFindAllPersonsQueryHandler(findAllPersons)
		bus.Register(getall.PersonQueryType, findAllPersonsQueryHandler)
	*/
}
