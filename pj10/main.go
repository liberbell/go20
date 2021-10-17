package main

import "fmt"

const Pi = 3.14

const (
	URL      = "http://www.yahoo.co.jp"
	Sitename = "Yahoo"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

const (
	c0 = iota
	c1
	c2
)

func main() {
	fmt.Println(Pi)

	Pi := 3
	fmt.Println(Pi)

	fmt.Println(URL, Sitename)

	fmt.Println(A, B, C, D, E, F)

	// var Big int = 9223372036854775807 + 1
	const Big = 9223372036854775807 + 1

	fmt.Println(Big - 1)

	fmt.Println(c0, c1, c2)

}
