package main

import "fmt"

func reciever(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}

}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// fmt.Println(<-ch1)

}
