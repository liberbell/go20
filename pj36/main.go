package main

import "fmt"

func main() {
	ch1 := make(chan int)
	fmt.Println(<-ch1)
}
