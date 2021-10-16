package main

import "fmt"

func main() {
	byteA := []byte{72, 73}
	fmt.Println(byteA)
	fmt.Printf("%T\n", byteA[0])

	fmt.Println(string(byteA))

	c := []byte("HI")
	fmt.Printf("%T", c)
	fmt.Println(c)
}
