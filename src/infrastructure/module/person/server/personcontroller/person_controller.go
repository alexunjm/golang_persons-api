package personcontroller

import (
	"fmt"
	"golang_persons-api/src/application/module/person/personcommand/create"
	"golang_persons-api/src/application/module/person/personcommand/update"
	"golang_persons-api/src/application/module/person/personquery"
	"golang_persons-api/src/domain/person/command"
	exception "golang_persons-api/src/infrastructure/httperror"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	message = gin.H{
		"test": "ok",
	}
)

// PersonController struct handles person requests
type PersonController struct {
	commandBus command.Bus
}

// NewPersonController creates controller for person requests
func NewPersonController(commandBus command.Bus) *PersonController {
	return &PersonController{commandBus}
}

// GetPerson func gets person by id
func (ctrl *PersonController) GetPerson(c *gin.Context) {

	personquery := personquery.Findpersonquery{ID: c.Param("person_id")}

	sample := personquery.GetImmutable()
	fmt.Printf("%+v", sample)
	fmt.Printf("id: %v",
		sample.GetID(),
	)

	c.JSON(http.StatusOK, personquery)
}

// GetAllPersons func gets all persons
func (ctrl *PersonController) GetAllPersons(c *gin.Context) {

	c.JSON(http.StatusOK, message)
}

// CreatePerson func creates a new person
func (ctrl *PersonController) CreatePerson(c *gin.Context) {
	var createPersonCommand create.CreatePersonCommand
	if err := c.ShouldBindJSON(&createPersonCommand); err != nil {
		theErr := exception.NewUnprocessableEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	ctrl.commandBus.Dispatch(createPersonCommand)

	c.JSON(http.StatusCreated, createPersonCommand)
}

// UpdatePerson func updates a person
func (ctrl *PersonController) UpdatePerson(c *gin.Context) {
	var updatePersonCommand update.UpdatePersonCommand
	if err := c.ShouldBindJSON(&updatePersonCommand); err != nil {
		theErr := exception.NewUnprocessableEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	ctrl.commandBus.Dispatch(updatePersonCommand)

	c.JSON(http.StatusOK, updatePersonCommand)
}

// DeletePerson func deletes a person
func (ctrl *PersonController) DeletePerson(c *gin.Context) {
	c.JSON(http.StatusOK, message)
}
