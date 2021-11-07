package main

import (
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)

	log.SetFlags(log.LstdFlags)
	log.Println("A")

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("B")

	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("C")

	log.SetFlags(log.LstdFlags)
	log.SetPrefix("[LOG]")
	log.Println("E")
}
