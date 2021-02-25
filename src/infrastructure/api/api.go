package api

import (
	"golang_persons-api/src/infrastructure/config"
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
		log.Print(".env file not found")
	}
}

// Run func starts the gin gonic api
func Run() {
	// loads default config and env config
	config.LoadAppEnv()

	// define API routes
	routes()

	// run on configured port
	Router.Run(config.ApplicationConfig.GetAPIPort())
}
