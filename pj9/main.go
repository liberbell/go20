package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// var i int = 1
	// fmt.Printf("Type of i : %T\n", i)
	// fl64 := float64(i)
	// fmt.Println(fl64)
	// fmt.Printf("Type of fl64: %T\n", fl64)

	// i2 := int(fl64)
	// fmt.Println(i2)
	// fmt.Printf("Type of i2: %T\n", i2)

	// fl := 10.3
	// fmt.Printf("Type of fl : %T\n", fl)
	// fmt.Println(fl)

	// i3 := int(fl)
	// fmt.Printf("Type of i3 : %T\n", i3)
	// fmt.Println(i3)

	var s string = "A"
	fmt.Printf("Type of s : %T\n", s)
	i, err := strconv.Atoi(s)
	fmt.Println(i)
	log.Println(err)

	var i2 int = 200
	s2 := strconv.Itoa(i2)
	fmt.Println(s2)
	fmt.Printf("Type of i2 : %T\n", s2)

	var h string = "hello world"
	b := []byte(h)
	fmt.Println(b)
	fmt.Printf("Type of h and b : %T %T\n", h, b)

	h2 := string(b)
	fmt.Println(h2)
	fmt.Printf("Type of h2 and b : %T %T\n", h2, b)

}
