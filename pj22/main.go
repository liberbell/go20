package main

import "fmt"

func anything(a interface{}) {
	fmt.Println(a)
}

func main() {
	anything("aaa")
	anything(1)

	var x interface{} = 3
	i := x.(int)
	// fmt.Println(x + i)
	fmt.Println(i + 1)
}
