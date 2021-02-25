package api

import (
	"fmt"
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

// Run func starts the gin api
func Run() error {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("No handled error")
			log.Fatal(r)
		}
	}()

	var err error
	// loads default config and env config
	err = config.LoadAppEnv()
	if err != nil {
		return err
	}

	// define API routes
	routes()

	// run on configured port
	err = Router.Run(config.APIPort)
	if err != nil {
		return err
	}
	return nil
}
