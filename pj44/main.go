package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	m := map[int]User{
		1: {Name: "Poul", Age: 10},
		2: {Name: "Jhon", Age: 20},
		3: {Name: "George", Age: 30},
		4: {Name: "Ringo", Age: 40},
	}
	fmt.Println(m)

	m2 := map[User]string{
		{Name: "Poul", Age: 10}: "London",
		{Name: "Jhon", Age: 20}: "Manchester",
	}
	fmt.Println(m2)

	m3 := make(map[int]User)
	fmt.Println(m3)
}
