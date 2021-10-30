package main

import "fmt"

type T struct {
	User User
}

type User struct {
	Name string
	Age  int
}

func main() {
	t := T{User: User{Name: "Bob", Age: 20}}
	fmt.Println(t)

	fmt.Println(t.User)
	fmt.Println(t.User.Name)
}
