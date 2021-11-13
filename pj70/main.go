package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{5, 3, 2, 4, 8, 9, 1, 0, 11, 7}
	s := []string{"a", "z", "h", "k"}

	fmt.Println(i, s)

	sort.Ints(i)
	sort.Strings(s)
	fmt.Println(i, s)
}
