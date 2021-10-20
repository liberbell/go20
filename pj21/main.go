package main

import "fmt"

func main() {
	// n := 1
	// switch n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("some number")
	// }

	// switch n := 2; n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("some number")
	// }

	// n := 6
	// switch {
	// case n > 0 && n < 4:
	// 	fmt.Println("less than 4")
	// case n > 3 && n < 7:
	// 	fmt.Println("from 4 to 7")
	// }

	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	case n > 3 && n < 6:
		fmt.Println("from 4 to 6")
	default:
		fmt.Println("some number")
	}
}
