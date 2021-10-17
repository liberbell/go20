package main

import "fmt"

const Pi = 3.14

const (
	URL      = "http://www.yahoo.co.jp"
	Sitename = "Yahoo"
)

func main() {
	fmt.Println(Pi)

	Pi := 3
	fmt.Println(Pi)

	fmt.Println(URL, Sitename)

}
