package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Users []*User

func main() {
	user1 := User{Name: "Bob", Age: 10}
	user2 := User{Name: "Eric", Age: 20}
	user3 := User{Name: "George", Age: 30}
	user4 := User{Name: "Ringo", Age: 40}

	users := Users{}
	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3, &user4)
	fmt.Println(users)
	for _, u := range users {
		fmt.Println(*u)
	}
}
