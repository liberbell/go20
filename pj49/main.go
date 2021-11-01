package main

import (
	"fmt"
	. "fmt"
	f "fmt"
	"foo"
)

func appName() string {
	const AppName = "GOApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func main() {
	f.Println(foo.Max)
	// fmt.Println(foo.min)
	f.Println(foo.ReturnMin())
	Println(foo.Max)

	fmt.Println(appName())
}
