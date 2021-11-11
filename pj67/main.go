package main

import "fmt"

func main() {
	go func ()  {
		for i := 0; i < 100 ; i++ {
			fmt.Println("1st Go routing.")
		}
	}
}