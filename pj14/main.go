package main

import "fmt"

func Plus(x, y int) int {
	return x + y
}

func main() {
	sum1 := Plus(3, 10)
	fmt.Println(sum1)
}
