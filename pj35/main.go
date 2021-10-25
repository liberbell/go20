package main

func main() {
	var ch1 := chan int

	// var ch2 <-chan int
	// var ch3 chan<- int

	ch1 = make(chan, int)
	ch2 := make(chan, int)
	
	fmt.println(cap(ch1))
	fmt.println(cap(ch2))
}