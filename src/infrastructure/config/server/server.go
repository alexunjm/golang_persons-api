package server

import (
	"context"
	"database/sql"
	"golang_persons-api/src/application/inmemory"
	"golang_persons-api/src/application/module/person/personcommand/create"
	"golang_persons-api/src/domain/person/command"
	"golang_persons-api/src/infrastructure/config/env"
	"golang_persons-api/src/infrastructure/config/storageconfig/mysqlstorage"
	"golang_persons-api/src/infrastructure/module/person/server/personrouter"
	"net/http"

	"github.com/gin-gonic/gin"
)

// New creates a new server
func New(ctx context.Context, envVariables env.EnvVariables) Server {

	srv := Server{
		Router:     gin.Default(),
		sqlDb:      mysqlstorage.NewConnection(envVariables),
		commandBus: nil,
	}

	srv.registerRoutes()

	// run on configured port
	srv.Router.Run(envVariables.GetAPIPort())
	return srv
}

type Server struct {
	Router *gin.Engine
	sqlDb  *sql.DB

	// deps
	commandBus command.Bus
}

func (s *Server) registerRoutes() {

	// handle default index response
	s.Router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "it's working!!")
	})

	var commandBus command.Bus = inmemory.NewCommandBus()

	handler := create.NewCreatePersonCommandHandler()
	commandBus.Register(create.PersonCommandType, handler)

	personrouter.HandleRoutes(s.Router, commandBus)

}
