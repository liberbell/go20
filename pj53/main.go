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

	t2 := tiem.Date(2021, 11, 3, 10, 10, 10, 0, time.Local)
}
