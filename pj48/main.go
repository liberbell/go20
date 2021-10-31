package main

import "fmt"

type Stringer interface {
	String() string
}

type Point struct {
	A int
	B string
}

func main() {
	p := &Point{100, "ABC"}
	fmt.Println(p)
}
