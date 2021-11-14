package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, _ := http.Get("https://example.com")
	fmt.Println(res.StatusCode)
}
