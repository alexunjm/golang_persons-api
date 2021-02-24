package person_controller

import (
	"fmt"
	"golang_persons-api/src/application/person/person_command"
	"golang_persons-api/src/application/person/person_query"
	"golang_persons-api/src/infrastructure/exception"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	message = gin.H{
		"test": "ok",
	}
)

type PersonController struct {
}

// GetPerson func gets person by id
func GetPerson(c *gin.Context) {

	personQuery := person_query.FindPersonQuery{ID: c.Param("person_id")}

	sample := personQuery.GetImmutable()
	fmt.Printf("%+v", sample)
	fmt.Printf("id: %v",
		sample.GetID(),
	)
	// // dispatch command
	// services.CommandBus.dispatch(&personQuery)

	c.JSON(http.StatusOK, personQuery)
}

// GetAllPersons func gets all persons
func GetAllPersons(c *gin.Context) {

	c.JSON(http.StatusOK, message)
}

// CreatePerson func creates a new person
func CreatePerson(c *gin.Context) {
	var personCommand person_command.CreatePersonCommand
	if err := c.ShouldBindJSON(&personCommand); err != nil {
		theErr := exception.NewUnprocessableEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	// sample := personCommand.GetImmutable()
	// fmt.Printf("%+v", sample)
	// fmt.Printf("id: %v, firstname: %v, lastname: %v, age: %v",
	// 	sample.GetID(),
	// 	sample.GetFirstname(),
	// 	sample.GetLastname(),
	// 	sample.GetAge(),
	// )
	// // dispatch command
	// services.CommandBus.dispatch(&personCommand)

	c.JSON(http.StatusCreated, personCommand)
}

// UpdatePerson func updates a person
func UpdatePerson(c *gin.Context) {
	var personCommand person_command.UpdatePersonCommand
	if err := c.ShouldBindJSON(&personCommand); err != nil {
		theErr := exception.NewUnprocessableEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	// sample := personCommand.GetImmutable()
	// fmt.Printf("%+v", sample)

	c.JSON(http.StatusOK, personCommand)
}

// DeletePerson func deletes a person
func DeletePerson(c *gin.Context) {
	c.JSON(http.StatusOK, message)
}

func getMessageID(personIDParam string) (int64, exception.MessageErr) {

	personID, personErr := strconv.ParseInt(personIDParam, 10, 64)
	if personErr != nil {
		return 0, exception.NewBadRequestError("person id should be a number")
	}
	return personID, nil
}

//Since we are going for the Person id more than we, we extracted this functionality to a function so we can have a DRY code.
// func getPersonId(msgIdParam string) (int64, error_utils.PersonErr) {
// 	msgId, msgErr := strconv.ParseInt(msgIdParam, 10, 64)
// 	if msgErr != nil {
// 		return 0, error_utils.NewBadRequestError("Person id should be a number")
// 	}
// 	return msgId, nil
// }
