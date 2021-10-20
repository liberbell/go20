package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var s string = "100"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("I = %T\n", i)
	} else {
		log.Println(err)
	}

}
