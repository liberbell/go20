package main

import "fmt"

func main() {
	var x interface{}
	fmt.Println(x)
	fmt.Printf("Type : %T\n", x)

	x = 1
	fmt.Println(x)
	fmt.Printf("Type : %T\n", x)

	x = 1.1
	fmt.Println(x)
	fmt.Printf("Type : %T\n", x)

	x = "Hello"
	fmt.Println(x)
	fmt.Printf("Type : %T\n", x)
}
