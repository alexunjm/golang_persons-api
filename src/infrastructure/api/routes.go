package api

import (
	"golang_persons-api/src/infrastructure/person/person_router"

	"github.com/gin-gonic/gin"

	"net/http"
)

func routes() {

	// handle default index response
	Router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "working")
	})

	person_router.HandleRoutes(Router)
}
