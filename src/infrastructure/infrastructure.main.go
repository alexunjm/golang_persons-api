package infrastructure

import (
	"log"

	"github.com/gin-gonic/gin"
	goDotEnv "github.com/joho/godotenv"
)

var (
	// Router handle all of URLs and mapping related
	Router = gin.Default()
)

func init() {
	//loads values from .env into the system
	if err := goDotEnv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

// StartApp func starts the gin gonic api
func StartApp() {
	loadAppEnv()

	routes()

	Router.Run(":8080")
}
