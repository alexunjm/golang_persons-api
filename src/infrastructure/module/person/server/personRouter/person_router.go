package personrouter

import (
	"golang_persons-api/src/infrastructure/module/person/server/personController"

	"github.com/gin-gonic/gin"
)

func HandleRoutes(Router *gin.Engine) {
	Router.GET("/persons/:person_id", personController.GetPerson)
	Router.GET("/persons", personController.GetAllPersons)
	Router.POST("/persons", personController.CreatePerson)
	Router.PUT("/persons/:person_id", personController.UpdatePerson)
	Router.DELETE("/persons/:person_id", personController.DeletePerson)
}
