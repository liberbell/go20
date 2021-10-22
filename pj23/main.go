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
		for j := 1; j < 3; j++ {
			if j > 1 {
				continue
			}
			fmt.Println(i, j, i*j)
		}
		fmt.Println("Don`t operate 1")
	}
}
