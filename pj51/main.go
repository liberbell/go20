package main

import (
	"fmt"
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

	// _, err := os.Open("test.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])
	// fmt.Println(os.Args[2])
	// fmt.Println(os.Args[3])

	// fmt.Printf("length= %d\n", len(os.Args))
	// for i, v := range os.Args {
	// 	fmt.Println(i, v)
	// }
	// f, err := os.Open("test.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer f.Close()

	// f, _ := os.Create("test.txt")
	// f.Write([]byte("Hello\n"))

	// f.WriteAt([]byte("Golang"), 6)
	// f.Seek(0, os.SEEK_END)

	// f.WriteString("Yarh")

	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(n)
	fmt.Println(string(bs))

	bs2 := make([]byte, 128)
	nn, err := f.ReadAt(bs2, 10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(nn)
	fmt.Println(string(bs2))
}
