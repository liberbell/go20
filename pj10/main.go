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

func main() {
	fmt.Println(Pi)

	Pi := 3
	fmt.Println(Pi)

	fmt.Println(URL, Sitename)

	fmt.Println(A, B, C, D, E, F)

	var Big int = 9223372036854775807 + 1

	fmt.Println(Big)

}
