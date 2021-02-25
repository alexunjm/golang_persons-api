package api

import (
	"golang_persons-api/src/infrastructure/module/person/server/personrouter"

	"github.com/gin-gonic/gin"

	"net/http"
)

func routes() {

	// handle default index response
	Router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "it's working!!")
	})

	personrouter.HandleRoutes(Router)
}
