package main

import "database/sql"

var Db *sql.DB

func main() {
	Db, _ := sql.Open("sqlite3", "./exampl.sql")
}
