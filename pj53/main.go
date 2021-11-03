package main

import (
	"fmt"
	"time"
)

func main() {
	// f, err := os.OpenFile("foo.txt", os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer f.Close()

	// bs := make([]byte, 128)
	// n, err := f.Read(bs)
	// fmt.Println(n)
	// fmt.Println(string(bs))

	t := time.Now()
	fmt.Println(t)

	t2 := time.Date(2021, 11, 4, 10, 10, 10, 0, time.Local)
	fmt.Println(t2)
	fmt.Println(t2.Year())
	fmt.Println(t2.Month())
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
	fmt.Println(t.Nanosecond())
	fmt.Println(t.Zone())
}
