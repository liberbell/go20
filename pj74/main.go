package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, _ := http.Get("https://example.com")
	fmt.Println(res.StatusCode)
	fmt.Println(res.Proto)
	fmt.Println(res.Header["Date"])
}
