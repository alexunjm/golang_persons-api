package mysqlstorage

import (
	exception "golang_persons-api/src/infrastructure/httperror"
	"strconv"
)

// mysql env variables
const (
	EnvDriver   = "MYSQL_DRIVER"
	EnvUsername = "MYSQL_USERNAME"
	EnvPassword = "MYSQL_PASSWORD"
	EnvHost     = "MYSQL_HOST"
	EnvDatabase = "MYSQL_DATABASE"
	EnvPort     = "MYSQL_PORT"
)

func NewMySQLDataConnection(dataConnection map[string]string) MySQLDataConnection {
	return MySQLDataConnection{
		driver:   dataConnection["driver"],
		username: dataConnection["username"],
		password: dataConnection["password"],
		host:     dataConnection["host"],
		database: dataConnection["database"],
		port:     dataConnection["port"],
	}
}

// MySQLDataConnection struct for SQLDataConnection
type MySQLDataConnection struct {
	driver   string
	username string
	password string
	host     string
	database string
	port     string
}

// Driver gets drive string value
func (msdc MySQLDataConnection) Driver() string {
	return msdc.driver
}

// Username gets username string value
func (msdc MySQLDataConnection) Username() string {
	return msdc.username
}

// Password gets password string value
func (msdc MySQLDataConnection) Password() string {
	return msdc.password
}

// Host gets host string value
func (msdc MySQLDataConnection) Host() string {
	return msdc.host
}

// Database gets database string value
func (msdc MySQLDataConnection) Database() string {
	return msdc.database
}

// Port gets port string value
func (msdc MySQLDataConnection) Port() int {
	port, _ := msdc.intPort()
	return port
}

// Port gets port string value
func (msdc MySQLDataConnection) intPort() (int, error) {
	port, err := strconv.Atoi(msdc.port)
	if err != nil {
		return 0, exception.NewInternalServerError("Port env variable should be a number")
	}
	return port, nil
}
