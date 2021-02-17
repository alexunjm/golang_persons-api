package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping function handles ping request
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
