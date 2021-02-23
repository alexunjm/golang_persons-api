package person_controller

import (
	"golang_persons-api/src/infrastructure/exception"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var message = gin.H{
	"test": "ok",
}

// GetPerson func gets person by id
func GetPerson(c *gin.Context) {
	// personID, err := getMessageID(c.Param("person_id"))
	// if err != nil {
	// 	c.JSON(err.Status(), err)
	// 	return
	// }
	c.JSON(http.StatusOK, message)
}

// GetAllPersons func gets all persons
func GetAllPersons(c *gin.Context) {
	c.JSON(http.StatusOK, message)
}

// CreatePerson func creates a new person
func CreatePerson(c *gin.Context) {
	c.JSON(http.StatusOK, message)
}

// UpdatePerson func updates a person
func UpdatePerson(c *gin.Context) {
	c.JSON(http.StatusOK, message)
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
