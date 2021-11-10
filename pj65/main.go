package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)

	fs1 := re.FindString("AAAAABCXYZZZZZZ")
	fmt.Println(fs1)
}
