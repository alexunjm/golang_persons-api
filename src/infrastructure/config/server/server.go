package server

import (
	"context"
	"database/sql"
	"golang_persons-api/src/application/inmemory"
	"golang_persons-api/src/domain/person/command"
	"golang_persons-api/src/infrastructure/config/env"
	"golang_persons-api/src/infrastructure/config/server/dependencyinjection"
	"golang_persons-api/src/infrastructure/config/storageconfig/mysqlstorage"
	"golang_persons-api/src/infrastructure/module/person/server/personcontroller"
	"golang_persons-api/src/infrastructure/module/person/server/personrouter"
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

	srv.registerRoutes()

	return srv
}

// Start runs API at given APIPort
func (s *Server) Start(APIPort string) {
	// run on configured port
	s.router.Run(APIPort)
}

func (s *Server) registerRoutes() {

	// handle default index response
	s.router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "it's working!!")
	})

	commandBus := inmemory.NewCommandBus()
	dependencyinjection.RegisterCommandHandlers(commandBus)

	personController := personcontroller.NewPersonController(commandBus)
	personrouter.HandleRoutes(s.router, personController)

}
