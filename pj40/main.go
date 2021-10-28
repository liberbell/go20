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
	user2.Name = "Eric"
	user2.Age = 20
	fmt.Println("Name: ", user2.Name)
	fmt.Println("Age: ", user2.Age)

	user3 := User{Name: "Jhon", Age: 30}
	fmt.Println("Name: ", user3.Name)
	fmt.Println("Age: ", user3.Age)
}
