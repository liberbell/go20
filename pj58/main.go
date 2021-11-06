package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println("Display")
	// fmt.Print("Hello")
	// fmt.Println("Hello!")

	// fmt.Println("hello", "world")
	// fmt.Println("hello" + "world")

	// fmt.Printf("%s\n", "hello")
	// fmt.Printf("%#v\n", "hello")

	// s := fmt.Sprint("Hello")
	// s1 := fmt.Sprintf("%v\n", "Hello")
	// s2 := fmt.Sprintln("Hello")

	// fmt.Println(s)
	// fmt.Println(s1)
	// fmt.Println(s2)

	fmt.Fprint(os.Stdout, "hello")
	fmt.Fprintf(os.Stdout, "hello")
	fmt.Fprintln(os.Stdout, "hello")

	f, _ := os.Create("test.txt")
	defer f.Close()

	fmt.Fprintln(f, "Fprint")
}
