package main

import (
	"fmt"
	"strconv"

	// i3 := int(fl)
	// fmt.Printf("Type of i3 : %T\n", i3)
	// fmt.Println(i3)

	var s string = "100"
	fmt.Printf("Type of s : %T\n", s)
	i, _ := strconv.Atoi(s)
	fmt.Println(i)
	fmt.Printf("Type of conv i : %T\n", i)
}
