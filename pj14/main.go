package main

import "fmt"

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func main() {
	sum1 := Plus(3, 10)
	fmt.Println(sum1)

	num, amari := Div(5, 2)
	fmt.Println(num, amari)
}
