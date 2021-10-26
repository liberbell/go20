package main

import "fmt"

func main() {
	// var ch1 chan int

	// var ch2 <-chan int
	// var ch3 chan<- int

	// ch1 = make(chan int)
	// ch2 := make(chan int)

	// fmt.Println(cap(ch1))
	// fmt.Println(cap(ch2))

	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	ch3 <- 1
	fmt.Println(len(ch3))
	// fmt.Println(byte(ch3))

	ch3 <- 2
	ch3 <- 3
	fmt.Println(len(ch3))

	i := <-ch3
	fmt.Println("len first", len(ch3), i)
	i2 := <-ch3
	fmt.Println("len second", len(ch3), i2)
	i3 := <-ch3
	fmt.Println("len third", len(ch3) i3)
}
