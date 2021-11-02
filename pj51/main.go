package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Exit(1)
	// fmt.Println("Start")

	defer func() {
		fmt.Println("defer")
	}()
	os.Exit(0)
}
