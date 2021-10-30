package main

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
}
