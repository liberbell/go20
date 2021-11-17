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

	fmt.Println("a")
}