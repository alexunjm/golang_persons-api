package person_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var message = gin.H{
	"test": "ok",
}

func GetPerson(c *gin.Context) {
	c.JSON(http.StatusOK, message)
}
func GetAllPersons(c *gin.Context) {
	c.JSON(http.StatusOK, message)
}
func CreatePerson(c *gin.Context) {
	c.JSON(http.StatusOK, message)
}
func UpdatePerson(c *gin.Context) {
	c.JSON(http.StatusOK, message)
}
func DeletePerson(c *gin.Context) {
	c.JSON(http.StatusOK, message)
}

//Since we are going for the Person id more than we, we extracted this functionality to a function so we can have a DRY code.
// func getPersonId(msgIdParam string) (int64, error_utils.PersonErr) {
// 	msgId, msgErr := strconv.ParseInt(msgIdParam, 10, 64)
// 	if msgErr != nil {
// 		return 0, error_utils.NewBadRequestError("Person id should be a number")
// 	}
// 	return msgId, nil
// }
