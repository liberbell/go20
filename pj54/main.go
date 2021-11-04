package main

import "time"

func main() {
	t5 := time.Date(2021, 11, 4, 10, 10, 10, 0, time.Local)
	t6 := time.Now()
	d2 := t5.Sub(t6)
}
