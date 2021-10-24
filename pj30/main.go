package main

import "fmt"

func main() {
	// sl := []int{100, 200}
	// sl2 := sl

	// fmt.Println(sl)
	// fmt.Println(sl2)
	// sl2[0] = 500
	// fmt.Println(sl)
	// fmt.Println(sl2)

	// var i int = 10
	// i2 := i
	// i2 = 100
	// fmt.Println(i)
	// fmt.Println(i2)

	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5, 10)
	n := copy(sl2, sl)

	fmt.Println(sl)
	fmt.Println(sl2)
	fmt.Println(n)
}
