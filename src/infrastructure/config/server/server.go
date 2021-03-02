package server

import (
	"context"
	"database/sql"
	"golang_persons-api/src/application/dependencyinjection"
	"golang_persons-api/src/application/inmemory"
	"golang_persons-api/src/domain/person/command"
	"golang_persons-api/src/infrastructure/config/env"
	"golang_persons-api/src/infrastructure/config/storage/mysqlstorage"
	"golang_persons-api/src/infrastructure/module/person/personcontroller"
	"golang_persons-api/src/infrastructure/module/person/personrouter"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server handles routes, api starting (run) and persistence config
type Server struct {
	router *gin.Engine
	sqlDb  *sql.DB

	// deps
	commandBus command.Bus
}

// New creates a new server
func New(ctx context.Context, envVariables env.EnvVariables) Server {

	srv := Server{
		router:     gin.Default(),
		sqlDb:      mysqlstorage.NewConnection(envVariables),
		commandBus: nil,
	}

	srv.registerIndexRoute()

	commandBus := inmemory.NewCommandBus()
	srv.registerControllers(commandBus)

	dependencyinjection.RegisterCommandHandlers(commandBus)

	return srv
}

// Start runs API at given APIPort
func (s *Server) Start(APIPort string) {
	// run on configured port
	s.router.Run(APIPort)
}

func (s *Server) registerIndexRoute() {

	// handle default index response
	s.router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "it's working!!")
	})

}

func (s *Server) registerControllers(commandBus command.Bus) {

	personController := personcontroller.NewPersonController(commandBus)
	personrouter.HandleRoutes(s.router, personController)

}
