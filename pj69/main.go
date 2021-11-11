package main

import "time"

type A struct{}

type User struct {
	ID      int
	Name    string
	Email   string
	created time.Time
	A       A
}

func main() {

	u := new(User)
	u.ID = 1
}
