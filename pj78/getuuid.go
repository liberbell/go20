package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuidobj, _ := uuid.NewUUID()
	fmt.Println("   ", uuidobj.String())

	uuidobj2, _ := uuid.NewRandom()
	fmt.Println("   ", uuidobj2)
}
