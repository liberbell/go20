package main

import (
	"fmt"
	"log"
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

	// t := time.Now()
	// fmt.Println(t)

	// t2 := time.Date(2021, 11, 4, 10, 10, 10, 0, time.Local)
	// fmt.Println(t2)
	// fmt.Println(t2.Year())
	// fmt.Println(t2.Month())
	// fmt.Println(t2.Hour())
	// fmt.Println(t2.Minute())
	// fmt.Println(t2.Second())
	// fmt.Println(t2.Nanosecond())
	// fmt.Println(t2.Zone())

	fmt.Println(time.Hour)
	fmt.Printf("%T\n", time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Nanosecond)
	fmt.Println(time.Microsecond)

	d, err := time.ParseDuration("2h30m")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(d)

	t3 := time.Now()
	t3 = t3.Add(2*time.Minute + 15*time.Second)
	fmt.Println(t3)
}
