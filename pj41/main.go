package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

func (u *User) SetName2(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "Bob"}
	user1.SayName()

	user1.SetName("Eric")
	user1.SayName()

	user1.SetName2("Jhon")
	user1.SayName()

	user2 := &User{Name: "ringo"}
	user2.SetName2("George")
	user2.SayName()
}
