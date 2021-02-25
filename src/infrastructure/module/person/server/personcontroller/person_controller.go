package personcontroller

import (
	"fmt"
	"golang_persons-api/src/application/person/personcommand"
	"golang_persons-api/src/application/person/personquery"
	exception "golang_persons-api/src/infrastructure/httperror"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	message = gin.H{
		"test": "ok",
	}
)

type personcontroller struct {
}

// GetPerson func gets person by id
func GetPerson(c *gin.Context) {

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
func GetAllPersons(c *gin.Context) {

	c.JSON(http.StatusOK, message)
}

// CreatePerson func creates a new person
func CreatePerson(c *gin.Context) {
	var personcommand personcommand.Createpersoncommand
	if err := c.ShouldBindJSON(&personcommand); err != nil {
		theErr := exception.NewUnprocessableEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	// sample := personcommand.GetImmutable()
	// fmt.Printf("%+v", sample)
	// fmt.Printf("id: %v, firstname: %v, lastname: %v, age: %v",
	// 	sample.GetID(),
	// 	sample.GetFirstname(),
	// 	sample.GetLastname(),
	// 	sample.GetAge(),
	// )
	// // dispatch command
	// services.CommandBus.dispatch(&personcommand)

	c.JSON(http.StatusCreated, personcommand)
}

// UpdatePerson func updates a person
func UpdatePerson(c *gin.Context) {
	var personcommand personcommand.Updatepersoncommand
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
