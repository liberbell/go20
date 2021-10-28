package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	var user1 User
	user1.Name = "Bob"
	user1.Age = 10
	fmt.Println("Name: ", user1.Name)
	fmt.Println("Age: ", user1.Age)

	user2 := User{}
	fmt.Println("Name: ", user2.Name)
	fmt.Println("Age: ", user2.Age)
}
