package main

import "fmt"

func main() {
	var s string = "hello golang"
	fmt.Println(s)
	fmt.Printf("Type of s: %T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("Type of s: %T\n", si)

	fmt.Println(`test
	test
		test`)

	fmt.Println("\"")
	fmt.Println(`"`)

	fmt.Println(string(s[0]))
}
