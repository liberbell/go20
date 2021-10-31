package main

import "fmt"

type Myint int

func main() {
	var mi Myint
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)

	i := 100
	fmt.Println(mi + i)
}
