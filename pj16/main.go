package main

import "fmt"

func Callfunction(f func()) {
	f()
}

func main() {
	Callfunction(func() {
		fmt.Println("This is inside function.")
	})
}
