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

	x = [3]int{1, 2, 3}
	fmt.Println(x)
	fmt.Printf("Type : %T\n", x)

	x = 2
	fmt.Println(x + 3)
}
