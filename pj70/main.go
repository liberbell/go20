package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	Name  string
	Value int
}

func main() {
	// i := []int{5, 3, 2, 4, 8, 9, 1, 0, 11, 7}
	// s := []string{"a", "z", "h", "k"}

	// fmt.Println(i, s)

	// sort.Ints(i)
	// sort.Strings(s)
	// fmt.Println(i, s)
	el := []Entry{
		{"A", 20},
		{"F", 40},
		{"i", 30},
		{"b", 10},
		{"t", 15},
		{"y", 30},
		{"c", 30},
		{"w", 30},
	}
	fmt.Println(el)

	sort.Slice(el, func(i, j int) bool { return el[i].Name < el[j].Name })
	fmt.Println("-----------")
	fmt.Println(el)
	fmt.Println("-----------")
}
