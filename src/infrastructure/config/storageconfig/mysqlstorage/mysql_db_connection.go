package mysqlstorage

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// // ...

// MysqlDBConnection db connection handler for mysql
type MysqlDBConnection struct {
}

func (mdbc MysqlDBConnection) GetDB(sqlURI string) *sql.DB {

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
