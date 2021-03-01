package personrouter

import (
	"golang_persons-api/src/domain/person/command"
	"golang_persons-api/src/infrastructure/module/person/server/personcontroller"

	"github.com/gin-gonic/gin"
)

func HandleRoutes(Router *gin.Engine, commandBus command.Bus) {

	ctrl := personcontroller.NewPersonController(commandBus)

	Router.GET("/persons/:person_id", ctrl.GetPerson)
	Router.GET("/persons", ctrl.GetAllPersons)
	Router.POST("/persons", ctrl.CreatePerson)
	Router.PUT("/persons/:person_id", ctrl.UpdatePerson)
	Router.DELETE("/persons/:person_id", ctrl.DeletePerson)
}
