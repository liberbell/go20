package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		max int
		msg string
		x   bool
	)
	flag.IntVar(&max, "n", 32, "Max use")
	flag.StringVar(&msg, "m", "", "Message")
	flag.BoolVar(&x, "x", false, "Extra option")

	flag.Parse()

	fmt.Println("Max number: ", max)
	fmt.Println("Message: ", msg)
	fmt.Println("Extra Option: ", x)
}
