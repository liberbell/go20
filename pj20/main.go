package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		fmt.Printf("Loop %s\n", i)
		time.Sleep(1)
		i++
	}
}
