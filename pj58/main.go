package main

import "fmt"

func main() {
	// fmt.Println("Display")
	// fmt.Print("Hello")
	// fmt.Println("Hello!")

	// fmt.Println("hello", "world")
	// fmt.Println("hello" + "world")

	// fmt.Printf("%s\n", "hello")
	// fmt.Printf("%#v\n", "hello")

	s := fmt.Sprint("Hello")
	s1 := fmt.Sprintf("Hello")
	s2 := fmt.Sprintln("Hello")

	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)
}
