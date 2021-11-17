package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, _ := sql.Open("postgres", "user=test_user1 dbname=testdb1 password=user1pass sslmode=disable")
	defer Db.Close()

	fmt.Println("a")
}
