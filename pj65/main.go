package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)

	fs1 := re.FindString("AAAAABCXYZZZZZZ")
	fmt.Println(fs1)
	fs2 := re.FindAllString("ABCXYZABCXYZ", -1)
	fmt.Println(fs2)

	re1 := regexp.MustCompile(`(\d+)-(\d+)-(\d)`)
	s := `
		0120-123-456
		111-222-333
		090-1234-5678
		`
}
