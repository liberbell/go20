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

	// fmt.Println(strconv.FormatFloat(123.456, 'E', -1, 64))
	// fmt.Println(strconv.FormatFloat(123.456, 'e', 2, 64))

	// fmt.Println(strconv.FormatFloat(123.456, 'f', -1, 64))
	// fmt.Println(strconv.FormatFloat(123.456, 'f', 2, 64))

	// fmt.Println(strconv.FormatFloat(123.456, 'g', -1, 64))
	// fmt.Println(strconv.FormatFloat(123456789.123, 'f', -1, 64))

	// fmt.Println(strconv.FormatFloat(123.456, 'g', 4, 64))
	// fmt.Println(strconv.FormatFloat(123456789.123, 'G', 4, 64))

	bt1, _ := strconv.ParseBool("true")
	fmt.Printf("%v %T\n", bt1, bt1)

	bt2, _ := strconv.ParseBool("1")
	bt3, _ := strconv.ParseBool("t")
	bt4, _ := strconv.ParseBool("T")
	bt5, _ := strconv.ParseBool("TRUE")
	bt6, _ := strconv.ParseBool("True")
}
