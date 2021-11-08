package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("A", "ABC")
	fmt.Println(match)
}
