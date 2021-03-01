package personcontroller

import (
	"fmt"
	"golang_persons-api/src/application/person/personcommand"
	"golang_persons-api/src/application/person/personquery"
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
	// // dispatch command
	// services.CommandBus.dispatch(&personquery)

	c.JSON(http.StatusOK, personquery)
}

// GetAllPersons func gets all persons
func (ctrl *PersonController) GetAllPersons(c *gin.Context) {

	c.JSON(http.StatusOK, message)
}

// CreatePerson func creates a new person
func (ctrl *PersonController) CreatePerson(c *gin.Context) {
	var personcommand personcommand.Personcommand
	if err := c.ShouldBindJSON(&personcommand); err != nil {
		theErr := exception.NewUnprocessableEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	c.JSON(http.StatusCreated, personcommand)
}

// UpdatePerson func updates a person
func (ctrl *PersonController) UpdatePerson(c *gin.Context) {
	var personcommand personcommand.Personcommand
	if err := c.ShouldBindJSON(&personcommand); err != nil {
		theErr := exception.NewUnprocessableEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	// sample := personcommand.GetImmutable()
	// fmt.Printf("%+v", sample)

	c.JSON(http.StatusOK, personcommand)
}

// DeletePerson func deletes a person
func (ctrl *PersonController) DeletePerson(c *gin.Context) {
	c.JSON(http.StatusOK, message)
}

// func getMessageID(personIDParam string) (int64, exception.MessageErr) {

// 	personID, personErr := strconv.ParseInt(personIDParam, 10, 64)
// 	if personErr != nil {
// 		return 0, exception.NewBadRequestError("person id should be a number")
// 	}
// 	return personID, nil
// }
