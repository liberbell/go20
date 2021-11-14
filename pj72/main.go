package main

import (
	"context"
	"fmt"
	"time"
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("START")
	time.Sleep(2 * time.Second)
	fmt.Println("END")
	ch <- "RESULT"
}

func main() {
	ch := make(chan string)
}
