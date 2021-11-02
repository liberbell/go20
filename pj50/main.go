package main

import (
	"alib"
	"fmt"
)

func IsONe(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsONe(1))
	fmt.Println(IsONe(0))

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(alib.Average(s))
}
