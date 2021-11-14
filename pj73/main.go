package main

import "net/url"

func main() {
	u, _ := url.Parse("http://example.com/search?a=l&b=2#top")
}
