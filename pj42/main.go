package main

import "fmt"

type T struct {
	User User
}

type User struct {
	Name string
	Age  int
}

func (u *User) SetName() {
	u.Name = "Anybody"
}

func main() {
	t := T{User: User{Name: "Bob", Age: 20}}
	fmt.Println(t)

	fmt.Println(t.User)
	fmt.Println(t.User.Name)

	t.User.SetName()
	fmt.Println(t)
}
