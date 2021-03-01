package api

import (
	"context"
	"fmt"
	"golang_persons-api/src/infrastructure/config/env"
	"golang_persons-api/src/infrastructure/config/server"
	"log"

	goDotEnv "github.com/joho/godotenv"
)

func init() {
	//loads values from .env into the system
	if err := goDotEnv.Load(); err != nil {
		log.Print(".env file not found")
	}
}

// Run func starts the gin api
func Run() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("No handled error")
			log.Fatal(r)
		}
	}()

	// loads default config and env config
	envVariables := env.LoadAppEnv()

	server.New(context.Background(), envVariables)
}
