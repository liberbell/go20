package main

import "fmt"

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	// ints := integers()
	// fmt.Println(ints())
	// fmt.Println(ints())
	// fmt.Println(ints())

	// otherints := integers()
	// fmt.Println(otherints())
	// fmt.Println(otherints())
	// fmt.Println(otherints())
	// fmt.Println(otherints())

	a := 0
	if a == 1 {
		fmt.Println("Two")
	} else if a == 1 {
		fmt.Println("One")
	} else {
		fmt.Println("I don`t know.")
	}

	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}
}
