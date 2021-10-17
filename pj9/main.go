package main

import "fmt"

func main() {
	var i int = 1
	fmt.Printf("Type of i : %T\n", i)
	fl64 := float64(i)
	fmt.Println(fl64)
	fmt.Printf("Type of fl64: %T\n", fl64)

	i2 := int(fl64)
	fmt.Println(i2)
	fmt.Printf("Type of i2: %T\n", i2)
}
