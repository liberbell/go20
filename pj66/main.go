package main

import (
	"fmt"
	"time"
)

var st struct {
	A, B, C int
}

func UpdateAndPrint(n int) {
	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st)
}

func main() {
	for i := 0; i < 5; i++ {
		UpdateAndPrint(i)
	}
}
