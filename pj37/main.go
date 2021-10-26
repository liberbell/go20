package main

import "fmt"

func reveiver(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + " END")
}

func main() {
	ch1 := make(chan int, 2)

	// ch1 <- 1
	// close(ch1)
	// ch1 <- 2

	// i := <-ch1
	// fmt.Println(i)
	// ch1 <- 1

	// i, ok := <-ch1
	// fmt.Println(i, ok)
	// i2, ok := <-ch1
	// fmt.Println(i2, ok)

	go reveiver("A", ch1)
	go reveiver("B", ch1)
	go reveiver("C", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	close(ch1)
}
