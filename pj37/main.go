package main

import "fmt"

func reveiver(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
	}
}

func main() {
	ch1 := make(chan int, 2)

	ch1 <- 1
	close(ch1)
	// ch1 <- 2

	// i := <-ch1
	// fmt.Println(i)
	// ch1 <- 1

	i, ok := <-ch1
	fmt.Println(i, ok)
	i2, ok := <-ch1
	fmt.Println(i2, ok)
}
