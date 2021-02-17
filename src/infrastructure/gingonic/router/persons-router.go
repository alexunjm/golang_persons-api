package router

import (
	"github.com/alexunjm/golang_persons-api/src/infrastructure/gingonic/controller"
)

type person struct {
	name     string
	lastname string
	age      int
}

// resource naming convention https://restfulapi.net/resource-naming/
func personsRouting() {

	router.GET("/persons", controller.GetAllPersons)
	router.POST("/persons", controller.CreatePerson)
	router.GET("/persons/:id", controller.GetPersonForGivenID)
	router.PUT("/persons/:id", controller.UpdatePersonForGivenID)
	router.DELETE("/persons/:id", controller.DeletePersonForGivenID)
}
