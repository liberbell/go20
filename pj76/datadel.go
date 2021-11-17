package main

import "database/sql"

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()
}
