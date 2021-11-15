package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello World.")
}

func main() {
	http.ListenAndServe(":8000", nil)
}
