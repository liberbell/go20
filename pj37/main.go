package main

func main() {
	ch1 := make(chan int, 2)
	close(ch1)
}
