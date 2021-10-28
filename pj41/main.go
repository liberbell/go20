package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

func main() {
	user1 := User{Name: "Bob"}
	user1.SayName()
}
