package router

import "github.com/gin-gonic/gin"

var router = gin.Default()

// Config function configs routing
func Config() {
	pingRouting()
	personsRouting()
	router.Run(":8080")
}
