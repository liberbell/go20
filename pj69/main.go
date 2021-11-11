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
	u.Name = "test"
	u.Email = "test@example.com"
	u.created = time.Now()
}
