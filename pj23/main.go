package main

import "fmt"

func main() {
	for {
		for {
			for {
				fmt.Println("START")
				break
			}
			fmt.Println("Don`t operate 1")
		}
		fmt.Println("Don`t operate 2")
	}
	fmt.Println("END")
}
