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

type PersonDependencyInjection struct {
	commandBus command.Bus
	queryBus   query.Bus
	db         *sql.DB
	dbTimeout  time.Duration
}

// NewPersonDependencyInjection initialize a PersonDependencyInjection
func NewPersonDependencyInjection(commandBus command.Bus, queryBus query.Bus, db *sql.DB, dbTimeout time.Duration) PersonDependencyInjection {
	return PersonDependencyInjection{
		commandBus, queryBus, db, dbTimeout,
	}
}

// RegisterPersonCommandHandlers func registers all command handlers
func (pdi PersonDependencyInjection) RegisterPersonCommandHandlers() {

	personRepository := repository.NewPersonRepository(pdi.db, pdi.dbTimeout)

	personCreator := create.NewPersonCreator(personRepository)
	createPersonCommandHandler := create.NewCreatePersonCommandHandler(personCreator)
	pdi.commandBus.Register(create.PersonCommandType, createPersonCommandHandler)

	personUpdater := update.NewPersonUpdater(personRepository)
	updatePersonCommandHandler := update.NewUpdatePersonCommandHandler(personUpdater)
	pdi.commandBus.Register(update.PersonCommandType, updatePersonCommandHandler)

	personDeleter := delete.NewPersonDeleter(personRepository)
	deletePersonCommandHandler := delete.NewDeletePersonCommandHandler(personDeleter)
	pdi.commandBus.Register(delete.PersonCommandType, deletePersonCommandHandler)
}

// RegisterPersonQueryHandlers func registers all query handlers
func (pdi PersonDependencyInjection) RegisterPersonQueryHandlers() {

	personRepository := repository.NewPersonRepository(pdi.db, pdi.dbTimeout)

	personFinder := getone.NewPersonFinder(personRepository)
	findPersonQueryHandler := getone.NewFindPersonQueryHandler(personFinder)
	pdi.queryBus.Register(getone.PersonQueryType, findPersonQueryHandler)
	/*
		// TODO: find all
		findAllPersons := getall.NewFindAllPersons(personRepository)
		findAllPersonsQueryHandler := getall.NewFindAllPersonsQueryHandler(findAllPersons)
		bus.Register(getall.PersonQueryType, findAllPersonsQueryHandler)
	*/
}
