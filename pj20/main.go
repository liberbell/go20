package main

import (
	"fmt"
)

func main() {
	// i := 0
	// for {
	// 	fmt.Printf("Loop %d\n", i)
	// 	time.Sleep(time.Millisecond * 10)
	// 	i++
	// 	if i == 100 {
	// 		break
	// 	}
	// }

	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}
}
