package main

import "fmt"

func Returnfunc() func() {
	return func() {
		fmt.Println("This is inside function.")
	}
}

func main() {
	// f := func(x, y int) int {
	// 	return x + y
	// }

	// i := f(1, 2)
	// fmt.Println(i)

	// i2 := func(x, y int) int {
	// 	return x + y
	// }(10, 20)
	// fmt.Println(i2)

	// f := Returnfunc()
	// f()
	Returnfunc()
}
