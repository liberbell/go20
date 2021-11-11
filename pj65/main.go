package main

import (
	"fmt"
	"regexp"
)

func main() {
	// re := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)

	// fs1 := re.FindString("AAAAABCXYZZZZZZ")
	// fmt.Println(fs1)
	// fs2 := re.FindAllString("ABCXYZABCXYZ", -1)
	// fmt.Println(fs2)

	// re1 := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	// s := `
	// 	0120-456-7899
	// 	111-222-333
	// 	090-1234-5678
	// 	`

	// ms := re1.FindAllStringSubmatch(s, -1)
	// fmt.Println(ms)

	// for _, v := range ms {
	// 	fmt.Println(v)
	// }

	// fmt.Println(ms[0][0])
	// fmt.Println(ms[0][1])
	// fmt.Println(ms[0][2])

	re2 := regexp.MustCompile(`\s+`)
	fmt.Println(re2.ReplaceAllString("AAA BBB CCC", ","))

	re3 := regexp.MustCompile(`、|。`)
	fmt.Println(re3.ReplaceAllString("今日は、いい天気です。", "-"))

	re4 := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)
	fmt.Println(re4.Split("ASHVJVHAKFABCXYZIDERPTEXABCXYZ)", -1))

	re5 := regexp.MustCompile(`\s+`)
	fmt.Println(re5.Split("aaaaaaaaaa     bbbbbb    ccccc", -1))
}
