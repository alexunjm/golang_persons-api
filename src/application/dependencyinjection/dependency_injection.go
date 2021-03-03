package dependencyinjection

import (
	"database/sql"
	"golang_persons-api/src/domain/module/person/command"
	"golang_persons-api/src/domain/module/person/query"
	"golang_persons-api/src/infrastructure/module/person/persondependencyinjection"
	"time"
)

// RegisterCommandHandlers func registers all command handlers
func RegisterCommandHandlers(bus command.Bus, db *sql.DB, dbTimeout time.Duration) {

	persondependencyinjection.RegisterPersonCommandHandlers(bus, db, dbTimeout)
}

// RegisterQueryHandlers func registers all query handlers
func RegisterQueryHandlers(bus query.Bus, db *sql.DB, dbTimeout time.Duration) {

	persondependencyinjection.RegisterPersonQueryHandlers(bus, db, dbTimeout)
}
