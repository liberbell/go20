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
	// ch1 <- 1
	// ch2 <- "B"
	// ch1 <- 2

	// v1 := <-ch1
	// v2 := <-ch2
	// fmt.Println(v1, v2)

	select {
	case v1 := <-ch1:
		fmt.Println(v1 + 1000)
	case v2 := <-ch2:
		fmt.Println(v2 + "!")
	default:
		fmt.Println("unknown")
	}

	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	go func() {
		for {
			i := <-ch3
			ch4 <- i * 2
		}
	}()

	go func() {
		for {
			i2 := <-ch4
			ch5 <- i2 - 1
		}
	}()

	n := 0
Loop1:
	for {
		select {
		case ch3 <- n:
			n++
		case i3 := <-ch5:
			fmt.Println("Receved : ", i3)
		default:
			if n > 100 {
				break Loop1
			}
		}
		// if n > 100 {
		// 	break
		// }
	}
}
