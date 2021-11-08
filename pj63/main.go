package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// // fmt.Println(scanner)

	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "Read error", err)
	// }

	f, _ := os.Open("test.txt")
	bs, _ := ioutil.ReadAll(f)
	fmt.Println(string(bs))

	if err := ioutil.WriteFile("bar.txt", bs, 0666); err != nil {
		log.Fatalln(err)
	}
}
