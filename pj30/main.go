package main

import "fmt"

func main() {
	sl := []int{100, 200}
	sl2 := sl

	fmt.Println(sl)
	fmt.Println(sl2)
	sl2[0] = 500
	fmt.Println(sl)
	fmt.Println(sl2)
}
