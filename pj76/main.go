package main

import "database/sql"

var Db *sql.DB

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	cmd := `CREATE TABLE IF NOT EXISTS person(
			name STRING,
			age INT)`
}
