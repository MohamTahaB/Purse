package dbmanager

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitiateConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:1234)/test")
	return db, err
}
