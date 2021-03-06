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

	// i1 := strings.Index("ABCDEFGHI", "A")
	// i2 := strings.Index("ABCDEFGHI", "D")
	// fmt.Println(i1, i2)

	// i3 := strings.LastIndex("ABCDABCD", "BC")
	// fmt.Println(i3)

	// i4 := strings.IndexAny("ABC", "ABC")

	// i5 := strings.LastIndexAny("ABC", "ABC")
	// fmt.Println(i4, i5)

	// b1 := strings.HasPrefix("ABC", "A")
	// b2 := strings.HasSuffix("ABC", "C")
	// fmt.Println(b1, b2)

	// b3 := strings.Contains("ABC", "B")
	// b4 := strings.ContainsAny("ABCDE", "BD")
	// fmt.Println(b3, b4)

	// i6 := strings.Count("ABCABC", "B")
	// i7 := strings.Count("ABCABC", "")
	// fmt.Println(i6, i7)

	// s3 := strings.Repeat("ABC", 4)
	// s4 := strings.Repeat("ABC", 0)
	// fmt.Println(s3, s4)

	// s5 := strings.Replace("AAAAAA", "A", "B", 1)
	// s6 := strings.Replace("AAAAAA", "A", "B", -1)
	// fmt.Println(s5, s6)

	// s7 := strings.Split("A,B,C,D,E,F", ",")
	// s8 := strings.SplitAfter("A,B,C,D,E,F", ",")
	// fmt.Println(s7)
	// fmt.Println(s8)

	// s9 := strings.SplitN("A,B,C,D,E,F", ",", 2)
	// fmt.Println(s9)
	// s10 := strings.SplitAfterN("A,B,C,D,E,F", ",", 4)
	// fmt.Println(s10)

	s11 := strings.ToLower("ABC")
	s12 := strings.ToUpper("abc")
	fmt.Println(s11, s12)

	s13 := strings.TrimSpace("    -     ABC    -     ")
	s14 := strings.TrimSpace("　　　　　　ABC　　　　　　")
	fmt.Println(s13)
	fmt.Println(s14)

	s15 := strings.Fields("a b c")
	fmt.Println(s15)
	fmt.Println(s15[1])

}
