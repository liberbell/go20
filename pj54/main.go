package main

import (
	"fmt"
	"time"
)

func main() {
	t5 := time.Date(2021, 11, 4, 10, 10, 10, 0, time.Local)
	t6 := time.Now()
	fmt.Println(t6)

	d2 := t5.Sub(t6)
	fmt.Println(d2)

	fmt.Println(t6.Before(t5))
	fmt.Println(t6.After(t5))
	fmt.Println(t5.Before(t6))
	fmt.Println(t5.After(t6))

}
