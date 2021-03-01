package mysqlstorage

import (
	"database/sql"
	"fmt"
	"golang_persons-api/src/infrastructure/config/env"

	_ "github.com/go-sql-driver/mysql"
)

func NewConnection(variables env.EnvVariables) *sql.DB {

	return MysqlDBConnection{
		username: variables.EnvMySQLUsername,
		password: variables.EnvMySQLPassword,
		host:     variables.EnvMySQLHost,
		port:     variables.GetMySQLPort(),
		database: variables.EnvDatabase,
	}.getDB()
}

// MysqlDBConnection db connection handler for mysql
type MysqlDBConnection struct {
	username string
	password string
	host     string
	port     int
	database string
}

// GetDB returns db connection
func (MySqlDBC MysqlDBConnection) getDB() *sql.DB {

	sqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		MySqlDBC.username,
		MySqlDBC.password,
		MySqlDBC.host,
		MySqlDBC.port,
		MySqlDBC.database,
	)

	db, err := sql.Open("mysql", sqlURI)
	if err != nil {
		panic(err)
	}
	// // See "Important settings" section.
	// db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)

	return db
}
