package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("foo.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
}
