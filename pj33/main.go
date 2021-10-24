package main

import "fmt"

func main() {
	var m = map[string]int{"A": 10, "B": 20, "C": 30}

	fmt.Println(m)

	m2 := map[string]int{"A": 10, "B": 20, "C": 30}
	fmt.Println(m2)
}
