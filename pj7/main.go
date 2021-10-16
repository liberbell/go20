package main

import "fmt"

func main() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("Type : %T\n", arr1)

	var arr2 [3]string = [3]string{"today", "is", "fine."}
	fmt.Println(arr2)
	fmt.Printf("Type : %T\n", arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)
	fmt.Printf("Type : %T\n", ar3)
}
