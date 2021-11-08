package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// fmt.Println(scanner)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
