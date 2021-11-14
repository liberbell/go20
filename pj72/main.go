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
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)

	defer cancel()
	go longProcess(ctx, ch)

	for {
		select {
		case <-ctx.Done():
		}
	}
}
