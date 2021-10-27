package main

import "fmt"

func main() {
	// ch1 := make(chan int, 3)
	// ch1 <- 1
	// ch1 <- 2
	// ch1 <- 3
	// close(ch1)

	// for i := range ch1 {
	// 	fmt.Println(i)
	// }

	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	// ch2 <- "A"

	// v1 := <-ch1
	// v2 := <-ch2
	// fmt.Println(v1, v2)

	select {
	case v1 := <-ch1:
		fmt.Println(v1 + 1000)
	case v2 := <-ch2:
		fmt.Println(v2 + "!")
	}
}
