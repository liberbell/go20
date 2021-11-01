package main

import "fmt"

func IsONe(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsONe(1))
}
