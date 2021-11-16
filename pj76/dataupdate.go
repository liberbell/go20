package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")

}
