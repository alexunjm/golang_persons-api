package router

import (
	"github.com/alexunjm/golang_persons-api/src/infrastructure/gingonic/controller"
)

func pingRouting() {

	router.GET("/ping", controller.Ping)
}
