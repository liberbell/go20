package main

import "database/sql"

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, _ := sql.Open("postgres", "user= dbname=test_db1 password=")
	defer Db.Close()
}
