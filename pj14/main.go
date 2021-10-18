package main

import (
	"fmt"
)

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func main() {
	sum1 := Plus(3, 10)
	fmt.Println(sum1)

	num, amari := Div(5, 2)
	fmt.Println(num, amari)

	p1 := Double(15)
	fmt.Println(p1)
}
