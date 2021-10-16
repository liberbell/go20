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
	fmt.Printf("Type : %T\n", arr3)

	arr4 := [...]string{"Tomorrow", "is", "another", "day."}
	fmt.Println(arr4)
	fmt.Printf("Type : %T\n", arr4)

	fmt.Println(arr2[1])

	arr2[2] = "bad"
	fmt.Println(arr2)

	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5)

	fmt.Println(len(arr1))
}
