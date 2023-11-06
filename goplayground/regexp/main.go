package main

import (
	"fmt"
	"regexp"
)

func main() {
	// golang中用字符串需要转义，因此golang需要用字符串表达正则表达式时，\本身需要转义，否则会导致unknown escape报错，或者用``包裹string
	strs := []string{"abc.abc.abc", "abcdabcdabc"}
	for _, str := range strs {
		matched, err := regexp.MatchString("^abc\\.abc\\..*$", str)
		fmt.Println(matched, err)
	}

	// for _, str := range strs {
	// 	matched, err := regexp.MatchString("^abc\.abc\..*$", str)
	// 	fmt.Println(matched, err)
	// }

	for _, str := range strs {
		matched, err := regexp.MatchString(`^abc\.abc\..*$`, str)
		fmt.Println(matched, err)
	}
}
