package main

import (
	"fmt"
)

func main() {
	// var i int = 100
	// fmt.Println(i)
	// fmt.Println(reflect.TypeOf(i))

	// var i2 int64 = 200
	// fmt.Println(reflect.TypeOf(i2))
	// fmt.Printf("Type of i2: %T\n", i2)
	// fmt.Printf("Convert type of i2: %T\n", int32(i2))
	// fmt.Println(i + 50)

	// fmt.Println(i + i2)
	// var fl64 float64 = 2.4
	// fmt.Println(fl64)
	// fmt.Printf("Type of fl64: %T\n", fl64)
	// var fl = 3.2
	// fmt.Println(fl)
	// fmt.Printf("Type of fl: %T\n", fl)

	// var fl32 float32 = 1.2
	// fmt.Println(fl32)
	// fmt.Printf("Type of fl32: %T\n", fl32)

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

}
