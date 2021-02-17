package gingonic

import "github.com/alexunjm/golang_persons-api/src/infrastructure/gingonic/router"

// StartServer function starts api
func StartServer() {
	// DebugMode : default
	// ReleaseMode
	// TestMode
	// gin.SetMode(gin.TestMode)
	router.Config()
}
