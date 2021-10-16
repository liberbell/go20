package main

import "fmt"

func main() {
	byteA := []byte{72, 73}
	fmt.Println(byteA)
	fmt.Printf("%T\n", byteA[0])

	fmt.Println(string(byteA))

	c := []byte("HI")
	fmt.Printf("Type is :%T\n", c)
	fmt.Println(c)
	fmt.Println(string(c))
}
