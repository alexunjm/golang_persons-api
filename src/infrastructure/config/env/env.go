package env

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

// PUBLIC
var (

	// DB connetion pointer
	DB *sql.DB

	// APIPort configured
	APIPort string
)

// LoadAppEnv load env variables, and connect to DB
func LoadAppEnv() EnvVariables {

	envVariables := EnvVariables{
		envAPIPort:       os.Getenv("APP_PORT"),
		EnvDriver:        os.Getenv("MYSQL_DRIVER"),
		EnvMySQLUsername: os.Getenv("MYSQL_USERNAME"),
		EnvMySQLPassword: os.Getenv("MYSQL_PASSWORD"),
		EnvMySQLHost:     os.Getenv("MYSQL_HOST"),
		envMySQLPort:     os.Getenv("MYSQL_PORT"),
		EnvDatabase:      os.Getenv("MYSQL_DATABASE"),
	}
	return envVariables
}

type EnvVariables struct {
	envAPIPort       string
	EnvDriver        string
	EnvMySQLUsername string
	EnvMySQLPassword string
	EnvMySQLHost     string
	envMySQLPort     string
	EnvDatabase      string
}

// GetAPIPort return API port
func (envVar EnvVariables) GetAPIPort() string {

	port, err := strconv.Atoi(envVar.envAPIPort)

	if err != nil {
		// return "", err
		panic(err)
	}

	return fmt.Sprintf(":%d", port)
}

// GetMySQLPort return MySQL port
func (envVar EnvVariables) GetMySQLPort() int {

	port, err := strconv.Atoi(envVar.envMySQLPort)

	if err != nil {
		panic(err)
	}

	return port
}
