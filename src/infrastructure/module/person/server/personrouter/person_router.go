package personrouter

import (
	"golang_persons-api/src/infrastructure/module/person/server/personcontroller"

	"github.com/gin-gonic/gin"
)

func HandleRoutes(Router *gin.Engine) {
	Router.GET("/persons/:person_id", personcontroller.GetPerson)
	Router.GET("/persons", personcontroller.GetAllPersons)
	Router.POST("/persons", personcontroller.CreatePerson)
	Router.PUT("/persons/:person_id", personcontroller.UpdatePerson)
	Router.DELETE("/persons/:person_id", personcontroller.DeletePerson)
}
