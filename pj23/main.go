package main

import "fmt"

func main() {
	// Loop:
	// 	for {
	// 		for {
	// 			for {
	// 				fmt.Println("START")
	// 				break Loop
	// 			}
	// 			fmt.Println("Don`t operate 1")
	// 		}
	// 		fmt.Println("Don`t operate 2")
	// 	}
	// 	fmt.Println("END")

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
		}
	}
}
