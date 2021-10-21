package main

import "fmt"

func main() {
	for {
		for {
			for {
				fmt.Println("START")
				break
			}
			fmt.Println("Don`t operate")
		}
		fmt.Println("Don`t operate")
	}
	fmt.Println("END")
}
