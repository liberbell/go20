package main

import (
	"fmt"
	"strconv"
)

func main() {
	// bt := true
	// fmt.Printf("%T\n", strconv.FormatBool(bt))

	// i := strconv.FormatInt(-100, 10)
	// fmt.Printf("%v, %T\n", i, i)

	// i2 := strconv.Itoa(100)
	// fmt.Printf("%v, %T\n", i2, i2)

	fmt.Println(strconv.FormatFloat(123.456, 'E', -1, 64))
	fmt.Println(strconv.FormatFloat(123.456, 'e', 2, 64))

	fmt.Println(strconv.FormatFloat(123.456, 'f', -1, 64))
	fmt.Println(strconv.FormatFloat(123.456, 'f', 2, 64))
}
