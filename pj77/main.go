package main

import "database/sql"

func main() {
	Db, _ := sql.Open("postgres", "user= dbname=test_db1 password=")
	defer Db.Close()
}
