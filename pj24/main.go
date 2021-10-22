package main

import "fmt"

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}
func main() {
	TestDefer()

	defer func() {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()
}
