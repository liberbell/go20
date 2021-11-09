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

	// x+ 0回以上繰り返し
	// x* 1回以上繰り返し
	// x? 0回以上、1回以下
	// x{n, m} n回以上、m回以下繰り返し最小マッチ
	// x{n, } n回以上繰り返す
	// x{n} n回繰り返す
	// x*? 0回以上繰り返す、最小マッチ
	// x+? 1回以上繰り返す、最小マッチ
	// x?? 0回以上、1回以下繰り返す

	// re5 := regexp.MustCompile("a+b*")
	// fmt.Println(re5.MatchString("ab"))
	// fmt.Println(re5.MatchString("a"))
	// fmt.Println(re5.MatchString("aaaaaaa"))
	// fmt.Println(re5.MatchString("b"))

	re8 := regexp.MustCompile(`[XYZ]`)
	fmt.Println(re8.MatchString("Y"))

	re9 := regexp.MustCompile(`^[0-9A-Za-z_]{3}$`)
	fmt.Println(re9.MatchString("ABC"))
	fmt.Println(re9.MatchString("abcdefg"))

	re10 := regexp.MustCompile(`[^0-9A-Za-z_]`)
	fmt.Println(re10.MatchString("ABC"))
	fmt.Println(re10.MatchString("abcdefg"))
	fmt.Println(re10.MatchString("$"))

}
