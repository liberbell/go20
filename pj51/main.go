package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Exit(1)
	// fmt.Println("Start")

	// defer func() {
	// 	fmt.Println("defer")
	// }()
	// os.Exit(0)

	// _, err := os.Open("test.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])
}
