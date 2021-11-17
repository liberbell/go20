package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	// cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	// _, err := Db.Exec(cmd, "Bob", 47)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	_, err := Db.Exec(cmd, "Eric", 68)
	if err != nil {
		log.Fatalln(err)
	}
}
