package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, err = sql.Open("postgres", "user=test_user1 dbname=testdb1 password=user1pass sslmode=disable")
	if err != nil {
		log.Panicln(err)
	}
	defer Db.Close()

	// cmd := "INSERT INTO persons (name, age) VALUES ($1, $2)"
	// _, err := Db.Exec(cmd, "Bob", 47)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	cmd := "SELECT * FROM persons WHERE age = ?"
	row := Db.QueryRow(cmd, 1000)
	var p Person
	err := row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age)

	cmd1 := "SELECT * FROM persons"
	rows, _ := Db.Query(cmd1)
	defer rows.Close()

	var pp []Person

}
