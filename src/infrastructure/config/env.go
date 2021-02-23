package config

import (
	"os"
)

var app App

type App struct {
	MySLQEnv MySQL
}

type MySQL struct {
	driver   string
	username string
	password string
	host     string
	database string
	port     string
}

func LoadAppEnv() {
	app = App{
		MySLQEnv: MySQL{
			driver:   os.Getenv("DRIVER"),
			username: os.Getenv("USERNAME"),
			password: os.Getenv("PASSWORD"),
			host:     os.Getenv("HOST"),
			database: os.Getenv("DATABASE"),
			port:     os.Getenv("PORT"),
		},
	}

}
