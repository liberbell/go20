package main

import (
	"log"
	"os"
)

func main() {
	// log.SetOutput(os.Stdout)

	// log.Print("Log\n")
	// log.Printf("Log%d\n", 3)
	// log.Println("Log2")

	// log.Fatal("Log\n")
	// log.Fatalln("Log2")
	// log.Fatalf("Log%d\n", 3)

	// log.Panic("Log\n")
	// log.Panicln("Log2")
	// log.Panicf("Log%d\n", 3)

	f, err := os.Create("err.txt")
	if err != nil {
		return
	}
	log.SetOutput(f)
	log.Println("write file")
}
