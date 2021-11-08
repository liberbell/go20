package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("A", "ABC")
	fmt.Println(match)

	re1, _ := regexp.Compile("A")
	match = re1.MatchString("ABC")
	fmt.Println(match)
}
