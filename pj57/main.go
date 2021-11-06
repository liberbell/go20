package main

import "flag"

func main() {
	var (
		max int
		msg string
		x   bool
	)
	flag.IntVar(&max, "n", 32, "Max use")
}
