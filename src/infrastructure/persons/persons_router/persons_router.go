package persons_router

import (
	"github.com/alexunjm/golang_persons-api/src/infrastructure"
	"github.com/alexunjm/golang_persons-api/src/infrastructure/persons/persons_controller"
)

func HandleRoutes() {
	infrastructure.Router.GET("/persons/:person_id", persons_controller.GetPerson)
	infrastructure.Router.GET("/persons", persons_controller.GetAllPersons)
	infrastructure.Router.POST("/persons", persons_controller.CreatePerson)
	infrastructure.Router.PUT("/persons/:person_id", persons_controller.UpdatePerson)
	infrastructure.Router.DELETE("/persons/:person_id", persons_controller.DeletePerson)
}
