package dependencyinjection

import (
	"golang_persons-api/src/domain/person/command"
	"golang_persons-api/src/infrastructure/module/person/persondependencyinjection"
)

// RegisterCommandHandlers func registers all command handlers
func RegisterCommandHandlers(bus command.Bus) {

	persondependencyinjection.RegisterPersonCommandHandlers(bus)
}
