package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var DB2 *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_golang")
	if err != nil {
		panic(err)
	}

	DB = db
}

func ConnectDB2() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_golang2")
	if err != nil {
		panic(err)
	}

	DB2 = db
}
