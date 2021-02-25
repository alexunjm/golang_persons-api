package config

import (
	"database/sql"
	"fmt"
	"golang_persons-api/src/infrastructure/config/storageconfig/mysqlstorage"
	exception "golang_persons-api/src/infrastructure/httperror"
	"strconv"
)

// LoadAppEnv load env variables, and connect to DB
func LoadAppEnv() error {

	applicationConfig = &app{
		appEnv: GetApplicationConfig(),
		sqlEnv: GetSQLDataConnection(),
	}

	var err error

	APIPort, err = applicationConfig.getAPIPort()
	if err != nil {
		return err
	}

	DB = applicationConfig.connect()
	return nil
}

type app struct {
	appEnv APPConfigContract
	sqlEnv SQLDataConnectionContract
}

// applicationConfig handles all infrastructure default config
var applicationConfig *app

// DB connetion pointer
var DB *sql.DB

// APIPort configured
var APIPort string

// getAPIPort gets configured api port
func (config *app) getAPIPort() (string, error) {
	port, err := strconv.Atoi(config.appEnv.Port())
	if err != nil {
		return "", exception.NewInternalServerError("Port env variable should be a number")
	}
	return fmt.Sprintf(":%d", port), nil
}

func (config *app) connect() *sql.DB {
	var sqlDataConn SQLDataConnectionContract
	sqlDataConn = config.sqlEnv

	sqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		sqlDataConn.Username(),
		sqlDataConn.Password(),
		sqlDataConn.Host(),
		sqlDataConn.Port(),
		sqlDataConn.Database(),
	)
	return mysqlstorage.MysqlDBConnection{}.GetDB(sqlURI)
}
