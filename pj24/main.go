package main

import "fmt"

func TestDefer() {
	defer fmt.Println("END")
}
func main() {
	a
}
