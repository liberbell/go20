package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuidobj, _ := uuid.NewUUID()
	fmt.Println("   ", uuidobj.String())
}
