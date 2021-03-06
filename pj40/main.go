package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func UserUpdate(user User) {
	user.Name = "nobody"
	user.Age = 99
}

func UserUpdate2(user *User) {
	user.Name = "nemo"
	user.Age = 999
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

	user4 := User{"Elton", 40}
	fmt.Println("Name: ", user4.Name)
	fmt.Println("Age: ", user4.Age)

	// user5 := User{50, "Sting"}
	// fmt.Println("Name: ", user5.Name)
	// fmt.Println("Age: ", user5.Age)

	user6 := User{Name: "Kate"}
	fmt.Println("Name: ", user6.Name)
	fmt.Println("Age: ", user6.Age)

	user7 := new(User)
	fmt.Println(user7)
	fmt.Println("Name: ", user7.Name)
	fmt.Println("Age: ", user7.Age)

	user8 := &User{}
	fmt.Println(user8)

	UserUpdate(user1)
	UserUpdate2(user8)

	fmt.Println(user1)
	fmt.Println(user8)

}
