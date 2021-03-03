package personrouter

import (
	"golang_persons-api/src/infrastructure/module/person/personcontroller"

	"github.com/gin-gonic/gin"
)

// HandleRoutes handles person routes
func HandleRoutes(router *gin.Engine, controller *personcontroller.PersonController) {

	router.GET("/persons/:person_id", controller.GetOnePerson)
	router.GET("/persons", controller.GetAllPersons)
	router.POST("/persons", controller.CreatePerson)
	router.PUT("/persons/:person_id", controller.UpdatePerson)
	router.DELETE("/persons/:person_id", controller.DeletePerson)
}
