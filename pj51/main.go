package main

import (
	"log"
	"os"
)

func main() {
	// os.Exit(1)
	// fmt.Println("Start")

	// defer func() {
	// 	fmt.Println("defer")
	// }()
	// os.Exit(0)

	_, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}
}
