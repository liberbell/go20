package main

import (
	f "fmt"
	"foo"
)

func main() {
	f.Println(foo.Max)
	// fmt.Println(foo.min)
	f.Println(foo.ReturnMin())
}
