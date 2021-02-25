package config

import (
	"golang_persons-api/src/infrastructure/config/storageconfig/mysqlstorage"
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

	dataConnection := mysqlstorage.NewMySQLDataConnection(map[string]string{
		"driver":   os.Getenv(mysqlstorage.EnvDriver),
		"username": os.Getenv(mysqlstorage.EnvUsername),
		"password": os.Getenv(mysqlstorage.EnvPassword),
		"host":     os.Getenv(mysqlstorage.EnvHost),
		"database": os.Getenv(mysqlstorage.EnvDatabase),
		"port":     os.Getenv(mysqlstorage.EnvPort),
	})
	return dataConnection
}
