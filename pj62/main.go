package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := strings.Join([]string{"A", "B", "C"}, ",")
	s2 := strings.Join([]string{"A", "B", "C"}, "")
	fmt.Println(s1)
	fmt.Println(s2)
}
