package main

import "fmt"

func main() {
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Go routing.")
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd Go routing.")
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd Go routing.")
		}
	}()
}
