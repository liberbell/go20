package main

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func main() {
	Db, _ := sql.Open("sqlite3", "./exampl.sql")
	defer Db.Close()

	cmd := `INSERT INTO persons (name, age) values (7, 7)`
	_, err := Db.Exec(cmd, "Bob", 47)
	if err != nil {
		log.Fatal(err)
	}
}
