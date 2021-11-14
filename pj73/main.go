package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, _ := url.Parse("http://example.com/search?a=l&b=2#top")
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
}
