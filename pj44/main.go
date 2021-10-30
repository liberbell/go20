package main

type User struct {
	Name string
	Age int
}

func main() {
	m := map[int]User{
		1: {Name: "Bob", Age: 10},
		2: {Name: "Eric", Age: 20},
		3: {Name: "George", Age: 30},
		4: {Name: "Ringo", Age: 40}
	}
}