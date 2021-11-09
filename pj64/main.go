package main

import (
	"fmt"
	"regexp"
)

func main() {
	// match, _ := regexp.MatchString("A", "ABC")
	// fmt.Println(match)

	// re1, _ := regexp.Compile("A")
	// match = re1.MatchString("ABC")
	// fmt.Println(match)

	// re2 := regexp.MustCompile("A")
	// match = re2.MatchString("ABC")
	// fmt.Println(match)

	// i 大文字小文字を意識しない
	// m マルチラインモード
	// s .が\nにマッチ
	// U 最小マッチへの変換　x*はx*?、x+はx+?に
	// re3 := regexp.MustCompile(`(?i)abc`)
	// match = re3.MatchString("ABC")
	// fmt.Println(match)

	// ^ 文頭(mフラグの場合行末も)
	// $ 文末(mフラグの場合行頭も)
	// /A 文頭
	// /z 文末
	// /b ASCII
	// /B 非ASCII

	// re4 := regexp.MustCompile(`^ABC$`)
	// match = re4.MatchString("ABC")
	// fmt.Println(match)

	// match = re4.MatchString("   ABC   ")
	// fmt.Println(match)

	re5 := regexp.MustCompile("a+b*")
	fmt.Println(re5.MatchString("ab"))
	fmt.Println(re5.MatchString("a"))
	fmt.Println(re5.MatchString("aaaaaaa"))
	fmt.Println(re5.MatchString("b"))

}
