package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	// close(ch1)
	// ch1 <- 2

	// i := <-ch1
	// fmt.Println(i)
	ch1 <- 1

	i, ok := <-ch1
	fmt.Println(i, ok)
}
