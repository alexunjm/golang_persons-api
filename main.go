package main

import (
	"golang_persons-api/src/infrastructure/api"
	"log"
)

func main() {

	if err := api.Run(); err != nil {
		log.Fatal(err)
	}
}
