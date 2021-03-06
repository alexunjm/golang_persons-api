package server

import (
	"context"
	"database/sql"
	"golang_persons-api/src/application/inmemory"
	"golang_persons-api/src/domain/module/person/command"
	"golang_persons-api/src/domain/module/person/query"
	"golang_persons-api/src/infrastructure/config/env"
	"golang_persons-api/src/infrastructure/config/storage/mysqlstorage"
	"golang_persons-api/src/infrastructure/module/person/personcontroller"
	"golang_persons-api/src/infrastructure/module/person/persondependencyinjection"
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

	aHalfMin, _ := time.ParseDuration("30s")
	personDI := persondependencyinjection.NewPersonDependencyInjection(
		srv.commandBus,
		srv.queryBus,
		srv.sqlDb,
		aHalfMin,
	)
	personDI.RegisterPersonCommandHandlers()
	personDI.RegisterPersonQueryHandlers()

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
