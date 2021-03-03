package api

import (
	"context"
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

	// loads default config and env config
	envVariables := env.LoadAppEnv()

	srv := server.New(context.Background(), envVariables)
	srv.Start(envVariables.GetAPIPort())
}
