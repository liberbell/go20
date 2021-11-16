package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func main() {
	Db, _ := sql.Open("sqlite3", "./exampl.sql")
	defer Db.Close()

	cmd := "INSERT INTO persons (name, age) values (?, ?)"
	_, err := Db.Exec(cmd, "Bob", 47)
	if err != nil {
		log.Fatal(err)
	}
}
