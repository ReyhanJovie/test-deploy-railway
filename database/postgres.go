package database

import (
	"database/sql"
	"fmt"
)

func ConnectDB(host, port, user, password, name string) (db *sql.DB, err error) {
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v name=%v sslmode=disable", host, port, user, password, name)

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}
	return
}
