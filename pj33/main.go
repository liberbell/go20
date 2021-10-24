package main

import "fmt"

func main() {
	var m = map[string]int{"A": 10, "B": 20, "C": 30}

	fmt.Println(m)

	m2 := map[string]int{"A": 10, "B": 20, "C": 30}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)

}
