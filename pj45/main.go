package main

import "fmt"

type Myint int

func main() {
	var mi Myint
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)
}
