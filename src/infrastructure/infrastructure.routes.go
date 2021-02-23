package infrastructure

import (
	"net/http"

	"github.com/alexunjm/golang_persons-api/src/infrastructure/persons/persons_router"

	"github.com/gin-gonic/gin"
)

func routes() {

	// handle default index response
	Router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "working")
	})

	persons_router.HandleRoutes()
}
