package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	var user1 User
	fmt.Println(user1.Name, user1.Age)
}
