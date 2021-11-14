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
	fmt.Println(res.Header["Content-Type"])

	fmt.Println(res.Request.Method)
	fmt.Println(res.Request.URL)
}
