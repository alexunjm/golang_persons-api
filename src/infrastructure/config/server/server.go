package server

import (
	"context"
	"database/sql"
	"golang_persons-api/src/application/dependencyinjection"
	"golang_persons-api/src/application/inmemory"
	"golang_persons-api/src/domain/module/person/command"
	"golang_persons-api/src/domain/module/person/query"
	"golang_persons-api/src/infrastructure/config/env"
	"golang_persons-api/src/infrastructure/config/storage/mysqlstorage"
	"golang_persons-api/src/infrastructure/module/person/personcontroller"
	"golang_persons-api/src/infrastructure/module/person/personrouter"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Server handles routes, api starting (run) and persistence config
type Server struct {
	router *gin.Engine
	sqlDb  *sql.DB

	// deps
	commandBus command.Bus
	queryBus   query.Bus
}

// New creates a new server
func New(ctx context.Context, envVariables env.EnvVariables) Server {

	srv := Server{
		router:     gin.Default(),
		sqlDb:      mysqlstorage.NewConnection(envVariables),
		commandBus: inmemory.NewCommandBus(),
		queryBus:   inmemory.NewQueryBus(),
	}

	srv.registerIndexRoute()

	srv.registerControllers()

	var duration time.Duration = 5000
	dependencyinjection.RegisterCommandHandlers(srv.commandBus, srv.sqlDb, duration)
	dependencyinjection.RegisterQueryHandlers(srv.queryBus, srv.sqlDb, duration)

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

func (s *Server) registerControllers() {

	personController := personcontroller.NewPersonController(s.commandBus, s.queryBus)
	personrouter.HandleRoutes(s.router, personController)

}
