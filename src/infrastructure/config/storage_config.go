package config

import (
	"golang_persons-api/src/infrastructure/config/storageConfig/mysqlStorage"
	"os"
)

// SQLDataConnectionContract interface for sql data connection
type SQLDataConnectionContract interface {
	Drive() string
	Username() string
	Password() string
	Host() string
	Database() string
	Port() string
}

// GetSQLDataConnection returns data connection for sql engine
func GetSQLDataConnection() SQLDataConnectionContract {

	return getMySQLDataConnection()
}

func getMySQLDataConnection() SQLDataConnectionContract {

	dataConnection := mysqlStorage.NewMySQLDataConnection(map[string]string{
		"driver":   os.Getenv(mysqlStorage.EnvDriver),
		"username": os.Getenv(mysqlStorage.EnvUsername),
		"password": os.Getenv(mysqlStorage.EnvPassword),
		"host":     os.Getenv(mysqlStorage.EnvHost),
		"database": os.Getenv(mysqlStorage.EnvDatabase),
		"port":     os.Getenv(mysqlStorage.EnvPort),
	})
	return dataConnection
}
