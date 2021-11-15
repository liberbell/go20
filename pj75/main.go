package main

import "net/http"

type MyHandler struct{}

func (h *MyHandler) ServeHTTP() {

}

func main() {
	http.ListenAndServe(":8000", nil)
}
