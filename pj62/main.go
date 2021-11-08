package main

import (
	"fmt"
	"strings"
)

func main() {
	// s1 := strings.Join([]string{"A", "B", "C"}, ",")
	// s2 := strings.Join([]string{"A", "B", "C"}, "")
	// fmt.Println(s1)
	// fmt.Println(s2)

	i1 := strings.Index("ABCDEFGHI", "A")
	i2 := strings.Index("ABCDEFGHI", "D")
	fmt.Println(i1, i2)

	i3 := strings.LastIndex("ABCDABCD", "BC")
	fmt.Println(i3)

	i4 := strings.IndexAny("ABC", "ABC")

	i5 := strings.LastIndexAny("ABC", "ABC")
	fmt.Println(i4, i5)

	b1 := strings.HasPrefix("ABC", "A")
	b2 := strings.HasSuffix("ABC", "C")
	fmt.Println(b1, b2)

	b3 := strings.Contains("ABC", "B")
	b4 := strings.ContainsAny("ABCDE", "BD")
	fmt.Println(b3, b4)

	i6 := strings.Count("ABCABC", "B")
	i7 := strings.Count("ABCABC", "")
	fmt.Println(i6, i7)
}
