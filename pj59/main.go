package main

import (
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)

	// log.Print("Log\n")
	// log.Printf("Log%d\n", 3)
	// log.Println("Log2")

	log.Fatal()
	log.Fatalln()
	log.Fatalf()
}
