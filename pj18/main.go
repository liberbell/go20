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
	if a == 2 {
		fmt.Println("Two")
	} else {
		fmt.Println("I don`t know.")
	}
}
