package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type person struct {
	name     string
	lastname string
	age      int
}

var persons = []person{
	{name: "Alex", lastname: "Jaramillo", age: 32},
	{name: "Andrea", lastname: "Laguna", age: 33},
	{name: "Jacob", lastname: "Jaramillo", age: 2},
}

// GetAllPersons function handles when get all persons
func GetAllPersons(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "all yours",
		"persons": persons,
	})
}

// CreatePerson function creates or add new person
func CreatePerson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":       "person has been created",
		"personCreated": persons[0],
	})
}

// GetPersonForGivenID function get existing person by id
func GetPersonForGivenID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":     "person exists",
		"personFound": persons[0],
	})
}

// UpdatePersonForGivenID function updates existing person by id
func UpdatePersonForGivenID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":       "person has been updated",
		"personUpdated": persons[0],
	})
}

// DeletePersonForGivenID function deletes existing person by id
func DeletePersonForGivenID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":       "person has been deleted",
		"personUpdated": persons[0],
	})
}
