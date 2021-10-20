package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		fmt.Printf("Loop %d\n", i)
		time.Sleep(time.Millisecond * 10)
		i++
	}
}
