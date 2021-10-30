package main

type T struct {
	User User
}

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func (u *User) SetName() {
	u.Name = "Anybody"
}

func main() {
	// t := T{User: User{Name: "Bob", Age: 20}}
	// fmt.Println(t)

	// fmt.Println(t.User)
	// fmt.Println(t.User.Name)

	// t.User.SetName()
	// fmt.Println(t)

}
