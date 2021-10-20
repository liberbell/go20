package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var s string = "ABC"

	i, err := strconv.Atoi(s)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("I = %T\n", i)
	}

}
