package person_router

import (
	"golang_persons-api/src/infrastructure/person/person_controller"

	"github.com/gin-gonic/gin"
)

func HandleRoutes(Router *gin.Engine) {
	Router.GET("/persons/:person_id", person_controller.GetPerson)
	Router.GET("/persons", person_controller.GetAllPersons)
	Router.POST("/persons", person_controller.CreatePerson)
	Router.PUT("/persons/:person_id", person_controller.UpdatePerson)
	Router.DELETE("/persons/:person_id", person_controller.DeletePerson)
}
