package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() *sql.DB {
	// Open database connection
	db, err := sql.Open("mysql", DbString)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	return db
}
