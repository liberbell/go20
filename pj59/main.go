package main

import (
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)

	log.Print("Log\n")
}
