package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 100
	fmt.Println(i)
	fmt.Println(reflect.TypeOf(i))

	var i2 int64 = 200
	fmt.Println(reflect.TypeOf(i2))
	fmt.Println(i + 50)

	// fmt.Println(i + i2)

}
