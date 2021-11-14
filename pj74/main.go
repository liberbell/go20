package main

import (
	"fmt"
	"net/url"
)

func main() {
	// res, _ := http.Get("https://example.com")
	// fmt.Println(res.StatusCode)
	// fmt.Println(res.Proto)
	// fmt.Println(res.Header["Date"])
	// fmt.Println(res.Header["Content-Type"])

	// fmt.Println(res.Request.Method)
	// fmt.Println(res.Request.URL)

	// defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(body))

	vs := url.Values{}
	vs.Add("id", "1")
	vs.Add("message", "Message")
	fmt.Println(vs.Encode())
}
